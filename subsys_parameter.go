package prarameters

// File       : subsys.go
// Path       : parameters
// Time       : CST 2023/4/26 11:24
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"encoding/json"
	"fmt"

	"github.com/Vanishvei/fass-sdk-parameters/utils"
)

var (
	defaultSectorSize   = 4096
	defaultShardingSize = "64G"
)

type ListSubsysParameter = listParameter

type RetrieveSubsysParameter struct {
	subsysName *string
}

func (parameter *RetrieveSubsysParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *RetrieveSubsysParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

func (parameter *RetrieveSubsysParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}

	return fmt.Sprintf("subsys/%s", *parameter.subsysName)
}

type subsysExportUnexportParameter struct {
	subsysName   *string
	protocolType *string
}

type ExportSubsysParameter subsysExportUnexportParameter

func (parameter *ExportSubsysParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

func (parameter *ExportSubsysParameter) GetProtocolType() string {
	return utils.StringValue(parameter.protocolType)
}

func (parameter *ExportSubsysParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf(fmt.Sprintf("subsys/%s/export", *parameter.subsysName))
}

func (parameter *ExportSubsysParameter) ExportISCSI() {
	parameter.protocolType = utils.String("iSCSI")
}

func (parameter *ExportSubsysParameter) ExportNVMeoF() {
	parameter.protocolType = utils.String("NVMeoF")
}

func (parameter *ExportSubsysParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*string{"protocol_type": parameter.protocolType})
}

func (parameter *ExportSubsysParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

type UnexportSubsysParameter subsysExportUnexportParameter

func (parameter *UnexportSubsysParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

func (parameter *UnexportSubsysParameter) GetProtocolType() string {
	return utils.StringValue(parameter.protocolType)
}

func (parameter *UnexportSubsysParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf(fmt.Sprintf("subsys/%s/unexport", *parameter.subsysName))
}

func (parameter *UnexportSubsysParameter) UnexportISCSI() {
	parameter.protocolType = utils.String("iSCSI")
}

func (parameter *UnexportSubsysParameter) UnexportNVMeoF() {
	parameter.protocolType = utils.String("NVMeoF")
}

func (parameter *UnexportSubsysParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*string{"protocol_type": parameter.protocolType})
}

func (parameter *UnexportSubsysParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

type DeleteSubsysParameter struct {
	subsysName     *string
	isForce        *bool
	isDeleteVolume *bool
}

func (parameter *DeleteSubsysParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

func (parameter *DeleteSubsysParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *DeleteSubsysParameter) ForceDelete() {
	parameter.isForce = utils.Bool(true)
}

func (parameter *DeleteSubsysParameter) DeleteVolume() {
	parameter.isDeleteVolume = utils.Bool(true)
}

func (parameter *DeleteSubsysParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s", *parameter.subsysName)
}

func (parameter *DeleteSubsysParameter) GetQuery() map[string]*string {
	_map := map[string]*string{}
	if parameter.isForce != nil {
		_map["is_force"] = utils.String("true")
	} else {
		_map["is_force"] = utils.String("false")
	}

	if parameter.isDeleteVolume != nil {
		_map["is_delete_volume"] = utils.String("true")
	} else {
		_map["is_delete_volume"] = utils.String("false")
	}
	return _map
}

type CreateSubsysParameter struct {
	bps           *int
	iops          *int
	capacity      *string
	format        *string
	inheritQos    *bool
	name          *string
	poolName      *string
	protocolType  *string
	sectorSize    *int
	sharding      *string
	shardingSize  *int
	snapshotName  *string
	srcVolumeName *string
	volumeName    *string
	enableISCSI   *bool
	enableNVMeoF  *bool
}

func (parameter *CreateSubsysParameter) MarshalJSON() ([]byte, error) {
	if parameter.enableISCSI != nil && parameter.enableNVMeoF != nil {
		parameter.protocolType = utils.String("all")
	} else if parameter.enableISCSI != nil {
		parameter.protocolType = utils.String("iSCSI")
	} else if parameter.enableNVMeoF != nil {
		parameter.protocolType = utils.String("NVMeoF")
	}

	if parameter.shardingSize != nil {
		_sharding := fmt.Sprintf("%dG", *parameter.shardingSize)
		parameter.sharding = utils.String(_sharding)
	}

	if parameter.name == nil {
		panic("parameter subsysName no set")
	}
	if parameter.poolName == nil {
		panic("parameter poolName no set")
	}
	if parameter.capacity == nil {
		panic("parameter capacity no set")
	}
	if parameter.protocolType == nil {
		panic("parameter protocolType no set")
	}

	if parameter.volumeName == nil {
		parameter.volumeName = parameter.name
	}

	_map := map[string]interface{}{
		"name":          *parameter.name,
		"capacity":      *parameter.capacity,
		"pool_name":     *parameter.poolName,
		"volume_name":   *parameter.volumeName,
		"protocol_type": *parameter.protocolType,
	}

	if parameter.bps != nil {
		_map["bps"] = *parameter.bps
	}

	if parameter.iops != nil {
		_map["iops"] = *parameter.iops
	}

	if parameter.format != nil {
		_map["format"] = *parameter.format
	}

	if parameter.sectorSize != nil {
		_map["sector_size"] = *parameter.sectorSize
	}

	if parameter.sharding != nil {
		_map["sharding"] = *parameter.sharding
	}

	return json.Marshal(_map)
}

func (parameter *CreateSubsysParameter) GetSubsysName() string {
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysParameter) GetVolumeName() string {
	if parameter.volumeName != nil {
		return utils.StringValue(parameter.volumeName)
	}
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysParameter) GetProtocolType() string {
	return utils.StringValue(parameter.protocolType)
}

func (parameter *CreateSubsysParameter) GetFormatType() string {
	return utils.StringValue(parameter.format)
}

func (parameter *CreateSubsysParameter) GetPoolName() string {
	return utils.StringValue(parameter.poolName)
}

func (parameter *CreateSubsysParameter) GetCapacity() string {
	return utils.StringValue(parameter.capacity)
}

func (parameter *CreateSubsysParameter) GetIOPS() int {
	return utils.IntValue(parameter.iops)
}

func (parameter *CreateSubsysParameter) GetBPS() int {
	return utils.IntValue(parameter.bps)
}

func (parameter *CreateSubsysParameter) GetSectorSize() int {
	if parameter.sectorSize == nil {
		return defaultSectorSize
	}
	return utils.IntValue(parameter.sectorSize)
}

func (parameter *CreateSubsysParameter) GetSharding() string {
	if parameter.shardingSize == nil {
		return defaultShardingSize
	}
	return utils.StringValue(parameter.sharding)
}

func (parameter *CreateSubsysParameter) SetName(subsysName string) {
	parameter.name = utils.String(subsysName)
}

func (parameter *CreateSubsysParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *CreateSubsysParameter) SetIops(iops int) {
	parameter.iops = utils.Int(iops)
}

func (parameter *CreateSubsysParameter) SetBpsMB(bps int) {
	parameter.bps = utils.Int(bps)
}

func (parameter *CreateSubsysParameter) SetCapacityGB(capacity int) {
	parameter.capacity = utils.String(fmt.Sprintf("%dg", capacity))
}

func (parameter *CreateSubsysParameter) SetCapacityTB(capacity int) {
	parameter.capacity = utils.String(fmt.Sprintf("%dt", capacity))
}

func (parameter *CreateSubsysParameter) SetFormatROW() {
	parameter.format = utils.String("row")
}

func (parameter *CreateSubsysParameter) SetFormatRAW() {
	parameter.format = utils.String("raw")
}

func (parameter *CreateSubsysParameter) SetPoolName(poolName string) {
	parameter.poolName = utils.String(poolName)
}

func (parameter *CreateSubsysParameter) EnableISCSI() {
	parameter.enableISCSI = utils.Bool(true)
}

func (parameter *CreateSubsysParameter) EnableNVMeoF() {
	parameter.enableNVMeoF = utils.Bool(true)
}

func (parameter *CreateSubsysParameter) SetSectorSize512() {
	parameter.sectorSize = utils.Int(512)
}

func (parameter *CreateSubsysParameter) SetSectorSize4096() {
	parameter.sectorSize = utils.Int(4096)
}

func (parameter *CreateSubsysParameter) SetSharding(sharding int) {
	parameter.shardingSize = utils.Int(sharding)
}

// ===========================================================================================================

type CreateSubsysFromSnapshotParameter CreateSubsysParameter

func (parameter *CreateSubsysFromSnapshotParameter) MarshalJSON() ([]byte, error) {
	if parameter.enableISCSI != nil && parameter.enableNVMeoF != nil {
		parameter.protocolType = utils.String("all")
	} else if parameter.enableISCSI != nil {
		parameter.protocolType = utils.String("iSCSI")
	} else if parameter.enableNVMeoF != nil {
		parameter.protocolType = utils.String("NVMeoF")
	}

	if parameter.shardingSize != nil {
		_sharding := fmt.Sprintf("%dG", *parameter.shardingSize)
		parameter.sharding = utils.String(_sharding)
	}

	if parameter.name == nil {
		panic("parameter subsysName no set")
	}
	if parameter.poolName == nil {
		panic("parameter poolName no set")
	}
	if parameter.protocolType == nil {
		panic("parameter protocolType no set")
	}
	if parameter.snapshotName == nil {
		panic("parameter srcSnapshotName no set")
	}
	if parameter.srcVolumeName == nil {
		panic("parameter srcVolumeName no set")
	}

	if parameter.volumeName == nil {
		parameter.volumeName = parameter.name
	}

	_map := map[string]interface{}{
		"name":            *parameter.name,
		"pool_name":       *parameter.poolName,
		"volume_name":     *parameter.volumeName,
		"protocol_type":   *parameter.protocolType,
		"snapshot_name":   *parameter.snapshotName,
		"src_volume_name": *parameter.srcVolumeName,
	}

	if parameter.bps != nil {
		_map["bps"] = *parameter.bps
	}

	if parameter.format != nil {
		_map["format"] = *parameter.format
	}

	if parameter.iops != nil {
		_map["iops"] = *parameter.iops
	}

	if parameter.inheritQos != nil {
		_map["inherit_qos"] = *parameter.inheritQos
	}

	if parameter.sectorSize != nil {
		_map["sector_size"] = *parameter.sectorSize
	}

	if parameter.sharding != nil {
		_map["sharding"] = *parameter.sharding
	}

	return json.Marshal(_map)
}

func (parameter *CreateSubsysFromSnapshotParameter) GetPoolName() string {
	return utils.StringValue(parameter.poolName)
}

func (parameter *CreateSubsysFromSnapshotParameter) GetSubsysName() string {
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysFromSnapshotParameter) GetVolumeName() string {
	if parameter.volumeName != nil {
		return utils.StringValue(parameter.volumeName)
	}
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysFromSnapshotParameter) GetSourceVolumeName() string {
	return utils.StringValue(parameter.srcVolumeName)
}

func (parameter *CreateSubsysFromSnapshotParameter) GetSourceSnapshotName() string {
	return utils.StringValue(parameter.snapshotName)
}

func (parameter *CreateSubsysFromSnapshotParameter) GetIOPS() int {
	return utils.IntValue(parameter.iops)
}

func (parameter *CreateSubsysFromSnapshotParameter) GetBPS() int {
	return utils.IntValue(parameter.bps)
}

func (parameter *CreateSubsysFromSnapshotParameter) GetSectorSize() int {
	if parameter.sectorSize == nil {
		return defaultSectorSize
	}
	return utils.IntValue(parameter.sectorSize)
}

func (parameter *CreateSubsysFromSnapshotParameter) GetSharding() string {
	if parameter.shardingSize == nil {
		return defaultShardingSize
	}
	return utils.StringValue(parameter.sharding)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetName(subsysName string) {
	parameter.name = utils.String(subsysName)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetPoolName(poolName string) {
	parameter.poolName = utils.String(poolName)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *CreateSubsysFromSnapshotParameter) InheritQos() {
	parameter.inheritQos = utils.Bool(true)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetBpsMB(bps int) {
	parameter.bps = utils.Int(bps)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetIops(iops int) {
	parameter.iops = utils.Int(iops)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetSharding(sharding int) {
	parameter.shardingSize = utils.Int(sharding)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetSrcVolumeName(srcVolumeName string) {
	parameter.srcVolumeName = utils.String(srcVolumeName)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetSrcSnapshotName(srcSnapshotName string) {
	parameter.snapshotName = utils.String(srcSnapshotName)
}

func (parameter *CreateSubsysFromSnapshotParameter) EnableISCSI() {
	parameter.enableISCSI = utils.Bool(true)
}

func (parameter *CreateSubsysFromSnapshotParameter) EnableNVMeoF() {
	parameter.enableNVMeoF = utils.Bool(true)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetSectorSize512() {
	parameter.sectorSize = utils.Int(512)
}

func (parameter *CreateSubsysFromSnapshotParameter) SetSectorSize4096() {
	parameter.sectorSize = utils.Int(4096)
}

// ==============================================================================================================

type CreateSubsysFromVolumeParameter CreateSubsysParameter

func (parameter *CreateSubsysFromVolumeParameter) MarshalJSON() ([]byte, error) {
	if parameter.enableISCSI != nil && parameter.enableNVMeoF != nil {
		parameter.protocolType = utils.String("all")
	} else if parameter.enableISCSI != nil {
		parameter.protocolType = utils.String("iSCSI")
	} else if parameter.enableNVMeoF != nil {
		parameter.protocolType = utils.String("NVMeoF")
	}

	if parameter.shardingSize != nil {
		_sharding := fmt.Sprintf("%dG", *parameter.shardingSize)
		parameter.sharding = utils.String(_sharding)
	}

	if parameter.name == nil {
		panic("parameter subsysName no set")
	}
	if parameter.poolName == nil {
		panic("parameter poolName no set")
	}
	if parameter.protocolType == nil {
		panic("parameter protocolType no set")
	}
	if parameter.srcVolumeName == nil {
		panic("parameter srcVolumeName no set")
	}

	if parameter.volumeName == nil {
		parameter.volumeName = parameter.name
	}

	_map := map[string]interface{}{
		"name":            *parameter.name,
		"pool_name":       *parameter.poolName,
		"volume_name":     *parameter.volumeName,
		"protocol_type":   *parameter.protocolType,
		"snapshot_name":   *parameter.name,
		"src_volume_name": *parameter.srcVolumeName,
	}

	if parameter.bps != nil {
		_map["bps"] = *parameter.bps
	}

	if parameter.format != nil {
		_map["format"] = *parameter.format
	}

	if parameter.iops != nil {
		_map["iops"] = *parameter.iops
	}

	if parameter.inheritQos != nil {
		_map["inherit_qos"] = *parameter.inheritQos
	}

	if parameter.sectorSize != nil {
		_map["sector_size"] = *parameter.sectorSize
	}

	if parameter.sharding != nil {
		_map["sharding"] = *parameter.sharding
	}

	return json.Marshal(_map)
}

func (parameter *CreateSubsysFromVolumeParameter) GetPoolName() string {
	return utils.StringValue(parameter.poolName)
}

func (parameter *CreateSubsysFromVolumeParameter) GetSubsysName() string {
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysFromVolumeParameter) GetVolumeName() string {
	if parameter.volumeName != nil {
		return utils.StringValue(parameter.volumeName)
	}
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysFromVolumeParameter) GetSourceVolumeName() string {
	return utils.StringValue(parameter.srcVolumeName)
}

func (parameter *CreateSubsysFromVolumeParameter) GetSourceSnapshotName() string {
	return utils.StringValue(parameter.snapshotName)
}

func (parameter *CreateSubsysFromVolumeParameter) GetProtocolType() string {
	return utils.StringValue(parameter.protocolType)
}

func (parameter *CreateSubsysFromVolumeParameter) GetIOPS() int {
	return utils.IntValue(parameter.iops)
}

func (parameter *CreateSubsysFromVolumeParameter) GetBPS() int {
	return utils.IntValue(parameter.bps)
}

func (parameter *CreateSubsysFromVolumeParameter) GetSharding() string {
	if parameter.shardingSize == nil {
		return defaultShardingSize
	}
	return utils.StringValue(parameter.sharding)
}

func (parameter *CreateSubsysFromVolumeParameter) SetName(subsysName string) {
	parameter.name = utils.String(subsysName)
}

func (parameter *CreateSubsysFromVolumeParameter) SetPoolName(poolName string) {
	parameter.poolName = utils.String(poolName)
}

func (parameter *CreateSubsysFromVolumeParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *CreateSubsysFromVolumeParameter) InheritQos() {
	parameter.inheritQos = utils.Bool(true)
}

func (parameter *CreateSubsysFromVolumeParameter) SetBpsMB(bps int) {
	parameter.bps = utils.Int(bps)
}

func (parameter *CreateSubsysFromVolumeParameter) SetIops(iops int) {
	parameter.iops = utils.Int(iops)
}

func (parameter *CreateSubsysFromVolumeParameter) SetSharding(sharding int) {
	parameter.shardingSize = utils.Int(sharding)
}

func (parameter *CreateSubsysFromVolumeParameter) SetSrcVolumeName(srcVolumeName string) {
	parameter.srcVolumeName = utils.String(srcVolumeName)
	parameter.snapshotName = parameter.name
}

func (parameter *CreateSubsysFromVolumeParameter) EnableISCSI() {
	parameter.enableISCSI = utils.Bool(true)
}

func (parameter *CreateSubsysFromVolumeParameter) EnableNVMeoF() {
	parameter.enableNVMeoF = utils.Bool(true)
}

func (parameter *CreateSubsysFromVolumeParameter) SetSectorSize512() {
	parameter.sectorSize = utils.Int(512)
}

func (parameter *CreateSubsysFromVolumeParameter) SetSectorSize4096() {
	parameter.sectorSize = utils.Int(4096)
}

// ==============================================================================================================

type RetrieveSubsysAuthParameter struct {
	subsysName *string
}

func (parameter *RetrieveSubsysAuthParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

func (parameter *RetrieveSubsysAuthParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *RetrieveSubsysAuthParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}

	return fmt.Sprintf("subsys/%s/auth", *parameter.subsysName)
}

type SetSubsysAuthParameter struct {
	subsysName *string
	groupName  *string
}

func (parameter *SetSubsysAuthParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

func (parameter *SetSubsysAuthParameter) GetGroupName() string {
	return utils.StringValue(parameter.groupName)
}

func (parameter *SetSubsysAuthParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"group_name": *parameter.groupName})
}

func (parameter *SetSubsysAuthParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *SetSubsysAuthParameter) SetGroupName(groupName string) {
	parameter.groupName = utils.String(groupName)
}

func (parameter *SetSubsysAuthParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/auth/set", *parameter.subsysName)
}

type RemoveSubsysAuthParameter struct {
	subsysName *string
}

func (parameter *RemoveSubsysAuthParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

func (parameter *RemoveSubsysAuthParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *RemoveSubsysAuthParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/auth/remove", *parameter.subsysName)
}

type RetrieveSubsysChapParameter struct {
	subsysName *string
}

func (parameter *RetrieveSubsysChapParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *RetrieveSubsysChapParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/chap", *parameter.subsysName)
}

type SetSubsysChapParameter struct {
	subsysName  *string
	accountName *string
}

func (parameter *SetSubsysChapParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"account_name": *parameter.accountName})
}

func (parameter *SetSubsysChapParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

func (parameter *SetSubsysChapParameter) GetAccountName() string {
	return utils.StringValue(parameter.accountName)
}

func (parameter *SetSubsysChapParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *SetSubsysChapParameter) SetAccountName(subsysName string) {
	parameter.accountName = utils.String(subsysName)
}

func (parameter *SetSubsysChapParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/chap/set", *parameter.subsysName)
}

type RemoveSubsysChapParameter struct {
	subsysName *string
}

func (parameter *RemoveSubsysChapParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

func (parameter *RemoveSubsysChapParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *RemoveSubsysChapParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/chap/remove", *parameter.subsysName)
}
