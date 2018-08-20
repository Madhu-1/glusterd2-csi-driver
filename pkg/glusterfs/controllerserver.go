package glusterfs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	csi "github.com/container-storage-interface/spec/lib/go/csi/v0"
	"github.com/gluster/gluster-csi-driver/pkg/glusterfs/utils"
	"github.com/gluster/glusterd2/pkg/api"
	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	glusterDescAnn            = "GlusterFS-CSI"
	glusterDescAnnValue       = "gluster.org/glusterfs-csi"
	defaultVolumeSize   int64 = 1000 * utils.MB // default volume size ie 1 GB
)

var volumeNotFound = errors.New("volume not found")

type ControllerServer struct {
	*GfDriver
}

func (cs *ControllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {

	var glusterServer string
	var bkpServers []string

	if req == nil {
		glog.Errorf("volume create request is nil")
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be empty")
	}

	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name is a required field")
	}
	glog.V(1).Infof("creating volume with name ", req.Name)

	if req.VolumeCapabilities == nil || len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume capabilities is a required field")
	}

	// If capacity mentioned, pick that or use default size 1 GB
	volSizeBytes := defaultVolumeSize
	if req.GetCapacityRange() != nil {
		volSizeBytes = int64(req.GetCapacityRange().GetRequiredBytes())
	}

	volSizeMB := int(utils.RoundUpSize(volSizeBytes, 1024*1024))

	// Get Volume name : TODO use the values from request
	volumeName := req.Name
	glusterVol := req.GetParameters()["glustervol"]
	glusterServer = req.GetParameters()["glusterserver"]
	glusterURL := req.GetParameters()["glusterurl"]
	glusterURLPort := req.GetParameters()["glusterurlport"]
	glusterUser := req.GetParameters()["glusteruser"]
	glusterUserSecret := req.GetParameters()["glusterusersecret"]

	glog.V(3).Infof("Request fields:[ %v %v %v %v %v %v]", glusterVol, glusterServer, glusterURL, glusterURLPort, glusterUser, glusterUserSecret)

	// Get list of volumes in the TSP
	volumeListResp, err := cs.client.Volumes("")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(volumeListResp) > 0 {
		glusterServer, bkpServers, err := cs.checkExistingVolumes(&volumeListResp, volumeName, volSizeBytes)
		if err != nil && err != volumeNotFound {
			return nil, err

		}
		if err == volumeNotFound {
			goto VOLUMECREATE
		}

		resp := &csi.CreateVolumeResponse{
			Volume: &csi.Volume{
				Id:            volumeName,
				CapacityBytes: int64(volSizeBytes),
				Attributes: map[string]string{
					"glustervol":        volumeName,
					"glusterserver":     glusterServer,
					"glusterbkpservers": strings.Join(bkpServers, ":"),
				},
			},
		}
		return resp, nil
	}
	glog.V(2).Infof("Volume list response: %+v", volumeListResp)

VOLUMECREATE:
	// If volume does not exist, provision volume
	glog.V(4).Infof("Received request to create/provision volume name:%s with size:%d", volumeName, volSizeMB)
	volMetaMap := make(map[string]string)
	volMetaMap[glusterDescAnn] = glusterDescAnnValue
	volumeReq := api.VolCreateReq{
		Name:     volumeName,
		Metadata: volMetaMap,
		Size:     uint64(volSizeMB),
	}

	glog.V(2).Infof("volume request: %+v", volumeReq)
	volumeCreateResp, err := cs.client.VolumeCreate(volumeReq)
	if err != nil {
		glog.Errorf("failed to create volume : %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create volume: %s", err.Error()))
	}

	glog.V(3).Infof("volume create response : %+v", volumeCreateResp)
	err = cs.client.VolumeStart(volumeName, true)
	if err != nil {
		//we dont need to delete the volume if volume start fails
		//as we are listing the volumes and starting it again
		//before sending back the response
		glog.Errorf("failed to start volume:%v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to start volume %s", err.Error()))
	}

	glusterServer, bkpServers, err = cs.getClusterNodes()

	if err != nil {
		glog.Errorf("failed to fetch details of cluster nodes: %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprintf("error in fecthing peer details %s", err.Error()))
	}

	resp := &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			Id:            volumeName,
			CapacityBytes: int64(volSizeBytes),
			Attributes: map[string]string{
				"glustervol":        volumeName,
				"glusterserver":     glusterServer,
				"glusterbkpservers": strings.Join(bkpServers, ":"),
			},
		},
	}

	glog.V(4).Infof("CSI Volume response: %+v", resp)
	return resp, nil
}

func (cs *ControllerServer) checkExistingVolumes(volumeListResp *api.VolumeListResp, volumeName string, volSizeBytes int64) (string, []string, error) {

	var tspServers []string
	var mountServer string
	var err error
	var found bool
	if volumeListResp == nil {
		glog.Errorf("Provided volume list response is nil")
		return "", nil, fmt.Errorf("volume list response is nil")
	}

	for _, vol := range *volumeListResp {
		if vol.Name == volumeName {
			found = true
			// Do the owner validation
			if glusterAnnVal, found := vol.Metadata[glusterDescAnn]; found {
				if glusterAnnVal != glusterDescAnnValue {
					return "", nil, status.Errorf(codes.Internal, "volume %s (%s) is not owned by Gluster CSI driver",
						vol.Name, vol.Metadata)
				}
			} else {
				return "", nil, status.Errorf(codes.Internal, "volume %s (%s) is not owned by Gluster CSI driver",
					vol.Name, vol.Metadata)
			}
			vsResp, e := cs.client.VolumeStatus(vol.Name)
			if e != nil {
				return "", nil, status.Errorf(codes.Internal, "failed to get volume status %s", e.Error())
			}
			if int64(vsResp.Size.Capacity) != volSizeBytes {
				return "", nil, status.Error(codes.AlreadyExists, fmt.Sprintf("volume already exits with different size: %d", vsResp.Size.Capacity))
			}

			//volume not started, start the volume
			if !vsResp.Online {
				err := cs.client.VolumeStart(vol.Name, true)
				if err != nil {
					return "", nil, status.Error(codes.Internal, fmt.Sprintf("failed to start volume"))
				}
			}

			glog.Info("Requested volume (%s) already exists in the storage pool", vol.Name)
			mountServer, tspServers, err = cs.getClusterNodes()

			if err != nil {
				return "", nil, status.Error(codes.Internal, fmt.Sprintf("error in fetching backup/peer server details %s", err.Error()))
			}

		}
	}
	if !found {
		return "", nil, volumeNotFound
	}
	return mountServer, tspServers, nil

}
func (cs *ControllerServer) getClusterNodes() (string, []string, error) {

	peers, err := cs.client.Peers()
	if err != nil {
		return "", nil, err
	}
	glusterServer := ""
	bkpservers := []string{}

	for i, p := range peers {
		if i == 0 {
			for _, a := range p.PeerAddresses {
				ip := strings.Split(a, ":")
				glusterServer = ip[0]
			}

			continue
		}
		for _, a := range p.PeerAddresses {
			ip := strings.Split(a, ":")
			bkpservers = append(bkpservers, ip[0])
		}

	}
	glog.V(2).Infof("Gluster server and Backup servers [%+v,%+v]", glusterServer, bkpservers)

	return glusterServer, bkpservers, err
}

// DeleteVolume deletes the given volume. The function is idempotent.
func (cs *ControllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	// TODO; return success if volume not found
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Volume delete request is nil")
	}

	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "volume ID is nil")
	}
	glog.V(2).Infof("Deleting volume with ID: %v", req.VolumeId)

	err := cs.client.VolumeStop(req.VolumeId)

	if err != nil {
		if cs.client.LastErrorResponse().StatusCode == http.StatusNotFound {
			return &csi.DeleteVolumeResponse{}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to stop volume %s", err.Error())
	}

	err = cs.client.VolumeDelete(req.VolumeId)
	if err != nil {
		if cs.client.LastErrorResponse().StatusCode == http.StatusNotFound {
			return &csi.DeleteVolumeResponse{}, nil
		}
		glog.Errorf("Volume delete failed :%v", err)
		return nil, status.Errorf(codes.Internal, "error deleting volume: %s", err.Error())
	}
	return &csi.DeleteVolumeResponse{}, nil
}

// ControllerPublishVolume attaches the given volume to the node
func (cs *ControllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

//ControllerUnpublishVolume deattaches the given volume from the node
func (cs *ControllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// ValidateVolumeCapabilities checks whether the volume capabilities requested
// are supported.
func (cs *ControllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "volume capabilities request is nil")
	}

	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "ValidateVolumeCapabilities() - Volume ID is nil")
	}

	if req.VolumeCapabilities == nil {
		return nil, status.Error(codes.InvalidArgument, "ValidateVolumeCapabilities is nil")
	}

	var vcaps []*csi.VolumeCapability_AccessMode
	for _, mode := range []csi.VolumeCapability_AccessMode_Mode{
		csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY,
		csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
	} {
		vcaps = append(vcaps, &csi.VolumeCapability_AccessMode{Mode: mode})
	}
	capSupport := true
	IsSupport := func(mode csi.VolumeCapability_AccessMode_Mode) bool {
		for _, m := range vcaps {
			if mode == m.Mode {
				return true
			}
		}
		return false
	}

	for _, cap := range req.VolumeCapabilities {
		if !IsSupport(cap.AccessMode.Mode) {
			capSupport = false
		}
	}
	resp := &csi.ValidateVolumeCapabilitiesResponse{
		Supported: capSupport,
	}
	glog.V(1).Infof("glusterfs CSI driver support capabilities: %v", resp)
	return resp, nil
}

// ListVolumes returns a list of all requested volumes
func (cs *ControllerServer) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {

	//Fetch all the volumes in the TSP
	volumes, err := cs.client.Volumes("")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var entries []*csi.ListVolumesResponse_Entry
	for _, vol := range volumes {
		v, e := cs.client.VolumeStatus(vol.Name)
		if e != nil {
			return nil, status.Errorf(codes.Internal, "failed to get volume status %s", err.Error())
		}
		entries = append(entries, &csi.ListVolumesResponse_Entry{
			Volume: &csi.Volume{
				Id:            vol.Name,
				CapacityBytes: (int64(v.Size.Capacity)) * utils.MB,
			},
		})
	}

	resp := &csi.ListVolumesResponse{
		Entries: entries,
	}

	return resp, nil

}

// GetCapacity returns the capacity of the storage pool
func (cs *ControllerServer) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {

	return nil, status.Error(codes.Unimplemented, "")
}

// ControllerGetCapabilities returns the capabilities of the controller service.
func (cs *ControllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	newCap := func(cap csi.ControllerServiceCapability_RPC_Type) *csi.ControllerServiceCapability {
		return &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: cap,
				},
			},
		}
	}

	var caps []*csi.ControllerServiceCapability
	for _, cap := range []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_LIST_VOLUMES,
	} {
		caps = append(caps, newCap(cap))
	}

	resp := &csi.ControllerGetCapabilitiesResponse{
		Capabilities: caps,
	}

	return resp, nil
}

//CreateSnapshot
func (cs *ControllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

//DeleteSnapshot
func (cs *ControllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

//ListSnapshots
func (cs *ControllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
