
<!---
This file is generated using commands described below. DO NOT EDIT.

$ curl -o endpoints.json -s -X GET http://127.0.0.1:24007/endpoints
$ go build pkg/tools/generate-doc.go
$ ./generate-doc
-->

# REST API Endpoints Reference

**Note:** Fields in request structs marked with "omitempty" struct tag are optional.

Name | Methods | Path | Request | Response
--- | --- | --- | --- | ---
GetVersion | GET | /version | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VersionResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VersionResp)
VolumeCreate | POST | /volumes | [VolCreateReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolCreateReq) | [VolumeCreateResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeCreateResp)
VolumeExpand | POST | /volumes/{volname}/expand | [VolExpandReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolExpandReq) | [VolumeExpandResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeExpandResp)
VolumeOptionGet | GET | /volumes/{volname}/options/{optname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeOptionGetResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeOptionGetResp)
VolumeOptionsGet | GET | /volumes/{volname}/options | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeOptionsGetResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeOptionsGetResp)
VolumeOptions | POST | /volumes/{volname}/options | [VolOptionReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolOptionReq) | [VolumeOptionResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeOptionResp)
VolumeReset | DELETE | /volumes/{volname}/options | [VolOptionResetReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolOptionResetReq) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
OptionGroupList | GET | /volumes/options-group | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [OptionGroupListResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#OptionGroupListResp)
OptionGroupCreate | POST | /volumes/options-group | [OptionGroupReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#OptionGroupReq) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
OptionGroupDelete | DELETE | /volumes/options-group/{groupname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolumeDelete | DELETE | /volumes/{volname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolumeInfo | GET | /volumes/{volname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeGetResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeGetResp)
VolumeBricksStatus | GET | /volumes/{volname}/bricks | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [BricksStatusResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#BricksStatusResp)
VolumeStatus | GET | /volumes/{volname}/status | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeStatusResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeStatusResp)
VolumeList | GET | /volumes | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeListResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeListResp)
VolumeStart | POST | /volumes/{volname}/start | [VolumeStartReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeStartReq) | [VolumeStartResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeStartResp)
VolumeStop | POST | /volumes/{volname}/stop | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeStopResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeStopResp)
Statedump | POST | /volumes/{volname}/statedump | [VolStatedumpReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolStatedumpReq) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolfilesGenerate | POST | /volfiles | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolfilesGet | GET | /volfiles | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolfilesGet | GET | /volfiles/{volfileid:.*} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
EditVolume | POST | /volumes/{volname}/edit | [VolEditReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolEditReq) | [VolumeEditResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeEditResp)
SnapshotCreate | POST | /snapshot | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotActivate | POST | /snapshot/{snapname}/activate | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotDeactivate | POST | /snapshot/{snapname}/deactivate | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotClone | POST | /snapshot/{snapname}/clone | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotRestore | POST | /snapshot/{snapname}/restore | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotInfo | GET | /snapshot/{snapname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotListAll | GET | /snapshots | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotStatus | GET | /snapshot/{snapname}/status | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotDelete | DELETE | /snapshot/{snapname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotConfigGet | GET | /snapshot/config | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotConfigSet | POST | /snapshot/config | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
SnapshotConfigReset | DELETE | /snapshot/config | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GetPeer | GET | /peers/{peerid} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [PeerGetResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerGetResp)
GetPeers | GET | /peers | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [PeerListResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerListResp)
DeletePeer | DELETE | /peers/{peerid} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
AddPeer | POST | /peers | [PeerAddReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerAddReq) | [PeerAddResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerAddResp)
EditPeer | POST | /peers/{peerid} | [PeerEditReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerEditReq) | [PeerEditResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerEditResp)
SetGlobalOptions | POST | /cluster/options | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GetGlobalOptions | GET | /cluster/options | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GeoReplicationCreate | POST | /geo-replication/{mastervolid}/{remotevolid} | [GeorepCreateReq](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepCreateReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSession)
GeoReplicationStart | POST | /geo-replication/{mastervolid}/{remotevolid}/start | [GeorepCommandsReq](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepCommandsReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSession)
GeoReplicationStop | POST | /geo-replication/{mastervolid}/{remotevolid}/stop | [GeorepCommandsReq](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepCommandsReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSession)
GeoReplicationDelete | DELETE | /geo-replication/{mastervolid}/{remotevolid} | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#)
GeoReplicationPause | POST | /geo-replication/{mastervolid}/{remotevolid}/pause | [GeorepCommandsReq](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepCommandsReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSession)
GeoReplicationResume | POST | /geo-replication/{mastervolid}/{remotevolid}/resume | [GeorepCommandsReq](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepCommandsReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSession)
GeoReplicationStatus | GET | /geo-replication/{mastervolid}/{remotevolid} | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSession)
GeoReplicationConfigGet | GET | /geo-replication/{mastervolid}/{remotevolid}/config | [GeorepOption](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepOption) | [GeorepOption](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepOption)
GeoReplicationConfigSet | POST | /geo-replication/{mastervolid}/{remotevolid}/config | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#)
GeoReplicationConfigReset | DELETE | /geo-replication/{mastervolid}/{remotevolid}/config | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#)
GeoReplicationStatusList | GET | /geo-replication | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSession)
GeoReplicationSshKeyGenerate | POST | /ssh-key/{volname}/generate | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#) | [GeorepSSHPublicKey](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSSHPublicKey)
GeoReplicationSshKeyPush | POST | /ssh-key/{volname}/push | [GeorepSSHPublicKey](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSSHPublicKey) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#)
GeoReplicationSshKeyGet | GET | /ssh-key/{volname} | [](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#) | [GeorepSSHPublicKey](https://godoc.org/github.com/gluster/glusterd2/plugins/georeplication/api/#GeorepSSHPublicKey)
BitrotEnable | POST | /volumes/{volname}/bitrot/enable | [](https://godoc.org/github.com/gluster/glusterd2/plugins/bitrot/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/bitrot/api/#)
BitrotDisable | POST | /volumes/{volname}/bitrot/disable | [](https://godoc.org/github.com/gluster/glusterd2/plugins/bitrot/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/bitrot/api/#)
BitrotScrubOndemand | POST | /volumes/{volname}/bitrot/scrubondemand | [](https://godoc.org/github.com/gluster/glusterd2/plugins/bitrot/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/bitrot/api/#)
BitrotScrubStatus | GET | /volumes/{volname}/bitrot/scrubstatus | [](https://godoc.org/github.com/gluster/glusterd2/plugins/bitrot/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/bitrot/api/#)
QuotaList | GET | /quota/{volname}/limit | [](https://godoc.org/github.com/gluster/glusterd2/plugins/quota/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/quota/api/#)
QuotaLimit | POST | /quota/{volname}/limit | [](https://godoc.org/github.com/gluster/glusterd2/plugins/quota/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/quota/api/#)
QuotaRemove | DELETE | /quota/{volname}/limit | [](https://godoc.org/github.com/gluster/glusterd2/plugins/quota/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/quota/api/#)
EventsWebhookAdd | POST | /events/webhook | [Webhook](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#Webhook) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#)
EventsWebhookTest | POST | /events/webhook/test | [Webhook](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#Webhook) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#)
EventsWebhookDelete | DELETE | /events/webhook | [WebhookDel](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#WebhookDel) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#)
EventsWebhookList | GET | /events/webhook | [](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#) | [WebhookList](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#WebhookList)
EventsList | GET | /events | [](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#) | [Event](https://godoc.org/github.com/gluster/glusterd2/plugins/events/api/#Event)
SelfHealInfo | GET | /volumes/{volname}/{opts}/heal-info | [](https://godoc.org/github.com/gluster/glusterd2/plugins/glustershd/api/#) | [BrickHealInfo](https://godoc.org/github.com/gluster/glusterd2/plugins/glustershd/api/#BrickHealInfo)
SelfHealInfo2 | GET | /volumes/{volname}/heal-info | [](https://godoc.org/github.com/gluster/glusterd2/plugins/glustershd/api/#) | [BrickHealInfo](https://godoc.org/github.com/gluster/glusterd2/plugins/glustershd/api/#BrickHealInfo)
SelfHeal | POST | /volumes/{volname}/heal | [](https://godoc.org/github.com/gluster/glusterd2/plugins/glustershd/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/glustershd/api/#)
DeviceAdd | POST | /devices/{peerid} | [AddDeviceReq](https://godoc.org/github.com/gluster/glusterd2/plugins/device/api/#AddDeviceReq) | [AddDeviceResp](https://godoc.org/github.com/gluster/glusterd2/plugins/device/api/#AddDeviceResp)
DeviceList | GET | /devices/{peerid} | [](https://godoc.org/github.com/gluster/glusterd2/plugins/device/api/#) | [ListDeviceResp](https://godoc.org/github.com/gluster/glusterd2/plugins/device/api/#ListDeviceResp)
DeviceEdit | POST | /devices/{peerid} | [EditDeviceReq](https://godoc.org/github.com/gluster/glusterd2/plugins/device/api/#EditDeviceReq) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/device/api/#)
DeviceListAll | GET | /devices | [](https://godoc.org/github.com/gluster/glusterd2/plugins/device/api/#) | [ListDeviceResp](https://godoc.org/github.com/gluster/glusterd2/plugins/device/api/#ListDeviceResp)
RebalanceStart | POST | /volumes/{volname}/rebalance/start | [StartReq](https://godoc.org/github.com/gluster/glusterd2/plugins/rebalance/api/#StartReq) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/rebalance/api/#)
RebalanceStop | POST | /volumes/{volname}/rebalance/stop | [](https://godoc.org/github.com/gluster/glusterd2/plugins/rebalance/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/rebalance/api/#)
RebalanceStatus | GET | /volumes/{volname}/rebalance | [](https://godoc.org/github.com/gluster/glusterd2/plugins/rebalance/api/#) | [](https://godoc.org/github.com/gluster/glusterd2/plugins/rebalance/api/#)
Statedump | GET | /statedump | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
List Endpoints | GET | /endpoints | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [ListEndpointsResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#ListEndpointsResp)
Glusterd2 service status | GET | /ping | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)