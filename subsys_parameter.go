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

type ListSubsys = listParameter

type subsysParameter struct {
	subsysName *string
}

func (parameter *subsysParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *subsysParameter) GetSubsysName() string {
	return utils.StringValue(parameter.subsysName)
}

type RetrieveSubsys struct {
	subsysParameter
}

func (parameter *RetrieveSubsys) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}

	return fmt.Sprintf("subsys/%s", *parameter.subsysName)
}

type subsysExportUnexportParameter struct {
	subsysParameter
	protocolType *string
}

func (parameter *subsysExportUnexportParameter) GetProtocolType() string {
	return utils.StringValue(parameter.protocolType)
}

func (parameter *subsysExportUnexportParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*string{"protocol_type": parameter.protocolType})
}

type ExportSubsys subsysExportUnexportParameter

func (parameter *ExportSubsys) ExportISCSI() {
	parameter.protocolType = utils.String("iSCSI")
}

func (parameter *ExportSubsys) ExportNVMeoF() {
	parameter.protocolType = utils.String("NVMeoF")
}

func (parameter *ExportSubsys) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf(fmt.Sprintf("subsys/%s/export", *parameter.subsysName))
}

type UnexportSubsys subsysExportUnexportParameter

func (parameter *UnexportSubsys) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf(fmt.Sprintf("subsys/%s/unexport", *parameter.subsysName))
}

func (parameter *UnexportSubsys) UnexportISCSI() {
	parameter.protocolType = utils.String("iSCSI")
}

func (parameter *UnexportSubsys) UnexportNVMeoF() {
	parameter.protocolType = utils.String("NVMeoF")
}

func (parameter *UnexportSubsys) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]*string{"protocol_type": parameter.protocolType})
}

type DeleteSubsys struct {
	subsysParameter
	isForce        *bool
	isDeleteVolume *bool
}

func (parameter *DeleteSubsys) ForceDelete() {
	parameter.isForce = utils.Bool(true)
}

func (parameter *DeleteSubsys) DeleteVolume() {
	parameter.isDeleteVolume = utils.Bool(true)
}

func (parameter *DeleteSubsys) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s", *parameter.subsysName)
}

func (parameter *DeleteSubsys) GetQuery() map[string]*string {
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

type CreateSubsys struct {
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

func (parameter *CreateSubsys) MarshalJSON() ([]byte, error) {
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

	if parameter.volumeName == nil {
		parameter.volumeName = parameter.name
	}

	_map := map[string]interface{}{
		"name":        *parameter.name,
		"capacity":    *parameter.capacity,
		"pool_name":   *parameter.poolName,
		"volume_name": *parameter.volumeName,
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

	if parameter.protocolType != nil {
		_map["protocol_type"] = *parameter.protocolType
	}

	return json.Marshal(_map)
}

func (parameter *CreateSubsys) GetSubsysName() string {
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsys) GetVolumeName() string {
	if parameter.volumeName != nil {
		return utils.StringValue(parameter.volumeName)
	}
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsys) GetProtocolType() string {
	return utils.StringValue(parameter.protocolType)
}

func (parameter *CreateSubsys) GetFormatType() string {
	return utils.StringValue(parameter.format)
}

func (parameter *CreateSubsys) GetPoolName() string {
	return utils.StringValue(parameter.poolName)
}

func (parameter *CreateSubsys) GetCapacity() string {
	return utils.StringValue(parameter.capacity)
}

func (parameter *CreateSubsys) GetIOPS() int {
	return utils.IntValue(parameter.iops)
}

func (parameter *CreateSubsys) GetBPS() int {
	return utils.IntValue(parameter.bps)
}

func (parameter *CreateSubsys) GetSectorSize() int {
	if parameter.sectorSize == nil {
		return defaultSectorSize
	}
	return utils.IntValue(parameter.sectorSize)
}

func (parameter *CreateSubsys) GetSharding() string {
	if parameter.shardingSize == nil {
		return defaultShardingSize
	}
	return utils.StringValue(parameter.sharding)
}

func (parameter *CreateSubsys) SetName(subsysName string) {
	parameter.name = utils.String(subsysName)
}

func (parameter *CreateSubsys) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *CreateSubsys) SetIops(iops int) {
	parameter.iops = utils.Int(iops)
}

func (parameter *CreateSubsys) SetBpsMB(bps int) {
	parameter.bps = utils.Int(bps)
}

func (parameter *CreateSubsys) SetCapacityGB(capacity int) {
	parameter.capacity = utils.String(fmt.Sprintf("%dg", capacity))
}

func (parameter *CreateSubsys) SetCapacityTB(capacity int) {
	parameter.capacity = utils.String(fmt.Sprintf("%dt", capacity))
}

func (parameter *CreateSubsys) SetFormatROW() {
	parameter.format = utils.String("row")
}

func (parameter *CreateSubsys) SetFormatRAW() {
	parameter.format = utils.String("raw")
}

func (parameter *CreateSubsys) SetPoolName(poolName string) {
	parameter.poolName = utils.String(poolName)
}

func (parameter *CreateSubsys) EnableISCSI() {
	parameter.enableISCSI = utils.Bool(true)
}

func (parameter *CreateSubsys) EnableNVMeoF() {
	parameter.enableNVMeoF = utils.Bool(true)
}

func (parameter *CreateSubsys) SetSectorSize512() {
	parameter.sectorSize = utils.Int(512)
}

func (parameter *CreateSubsys) SetSectorSize4096() {
	parameter.sectorSize = utils.Int(4096)
}

func (parameter *CreateSubsys) SetSharding(sharding int) {
	parameter.shardingSize = utils.Int(sharding)
}

// ===========================================================================================================

type CreateSubsysFromSnapshot CreateSubsys

func (parameter *CreateSubsysFromSnapshot) MarshalJSON() ([]byte, error) {
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

	if parameter.protocolType != nil {
		_map["protocol_type"] = *parameter.protocolType
	}

	return json.Marshal(_map)
}

func (parameter *CreateSubsysFromSnapshot) GetPoolName() string {
	return utils.StringValue(parameter.poolName)
}

func (parameter *CreateSubsysFromSnapshot) GetSubsysName() string {
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysFromSnapshot) GetVolumeName() string {
	if parameter.volumeName != nil {
		return utils.StringValue(parameter.volumeName)
	}
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysFromSnapshot) GetSourceVolumeName() string {
	return utils.StringValue(parameter.srcVolumeName)
}

func (parameter *CreateSubsysFromSnapshot) GetSourceSnapshotName() string {
	return utils.StringValue(parameter.snapshotName)
}

func (parameter *CreateSubsysFromSnapshot) GetIOPS() int {
	return utils.IntValue(parameter.iops)
}

func (parameter *CreateSubsysFromSnapshot) GetBPS() int {
	return utils.IntValue(parameter.bps)
}

func (parameter *CreateSubsysFromSnapshot) GetSectorSize() int {
	if parameter.sectorSize == nil {
		return defaultSectorSize
	}
	return utils.IntValue(parameter.sectorSize)
}

func (parameter *CreateSubsysFromSnapshot) GetSharding() string {
	if parameter.shardingSize == nil {
		return defaultShardingSize
	}
	return utils.StringValue(parameter.sharding)
}

func (parameter *CreateSubsysFromSnapshot) SetName(subsysName string) {
	parameter.name = utils.String(subsysName)
}

func (parameter *CreateSubsysFromSnapshot) SetPoolName(poolName string) {
	parameter.poolName = utils.String(poolName)
}

func (parameter *CreateSubsysFromSnapshot) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *CreateSubsysFromSnapshot) InheritQos() {
	parameter.inheritQos = utils.Bool(true)
}

func (parameter *CreateSubsysFromSnapshot) SetBpsMB(bps int) {
	parameter.bps = utils.Int(bps)
}

func (parameter *CreateSubsysFromSnapshot) SetIops(iops int) {
	parameter.iops = utils.Int(iops)
}

func (parameter *CreateSubsysFromSnapshot) SetSharding(sharding int) {
	parameter.shardingSize = utils.Int(sharding)
}

func (parameter *CreateSubsysFromSnapshot) SetSrcVolumeName(srcVolumeName string) {
	parameter.srcVolumeName = utils.String(srcVolumeName)
}

func (parameter *CreateSubsysFromSnapshot) SetSrcSnapshotName(srcSnapshotName string) {
	parameter.snapshotName = utils.String(srcSnapshotName)
}

func (parameter *CreateSubsysFromSnapshot) EnableISCSI() {
	parameter.enableISCSI = utils.Bool(true)
}

func (parameter *CreateSubsysFromSnapshot) EnableNVMeoF() {
	parameter.enableNVMeoF = utils.Bool(true)
}

func (parameter *CreateSubsysFromSnapshot) SetSectorSize512() {
	parameter.sectorSize = utils.Int(512)
}

func (parameter *CreateSubsysFromSnapshot) SetSectorSize4096() {
	parameter.sectorSize = utils.Int(4096)
}

// ==============================================================================================================

type CreateSubsysFromVolume CreateSubsys

func (parameter *CreateSubsysFromVolume) MarshalJSON() ([]byte, error) {
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

	if parameter.protocolType != nil {
		_map["protocol_type"] = *parameter.protocolType
	}

	if parameter.sharding != nil {
		_map["sharding"] = *parameter.sharding
	}

	return json.Marshal(_map)
}

func (parameter *CreateSubsysFromVolume) GetPoolName() string {
	return utils.StringValue(parameter.poolName)
}

func (parameter *CreateSubsysFromVolume) GetSubsysName() string {
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysFromVolume) GetVolumeName() string {
	if parameter.volumeName != nil {
		return utils.StringValue(parameter.volumeName)
	}
	return utils.StringValue(parameter.name)
}

func (parameter *CreateSubsysFromVolume) GetSourceVolumeName() string {
	return utils.StringValue(parameter.srcVolumeName)
}

func (parameter *CreateSubsysFromVolume) GetSourceSnapshotName() string {
	return utils.StringValue(parameter.snapshotName)
}

func (parameter *CreateSubsysFromVolume) GetProtocolType() string {
	return utils.StringValue(parameter.protocolType)
}

func (parameter *CreateSubsysFromVolume) GetIOPS() int {
	return utils.IntValue(parameter.iops)
}

func (parameter *CreateSubsysFromVolume) GetBPS() int {
	return utils.IntValue(parameter.bps)
}

func (parameter *CreateSubsysFromVolume) GetSharding() string {
	if parameter.shardingSize == nil {
		return defaultShardingSize
	}
	return utils.StringValue(parameter.sharding)
}

func (parameter *CreateSubsysFromVolume) SetName(subsysName string) {
	parameter.name = utils.String(subsysName)
}

func (parameter *CreateSubsysFromVolume) SetPoolName(poolName string) {
	parameter.poolName = utils.String(poolName)
}

func (parameter *CreateSubsysFromVolume) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *CreateSubsysFromVolume) InheritQos() {
	parameter.inheritQos = utils.Bool(true)
}

func (parameter *CreateSubsysFromVolume) SetBpsMB(bps int) {
	parameter.bps = utils.Int(bps)
}

func (parameter *CreateSubsysFromVolume) SetIops(iops int) {
	parameter.iops = utils.Int(iops)
}

func (parameter *CreateSubsysFromVolume) SetSharding(sharding int) {
	parameter.shardingSize = utils.Int(sharding)
}

func (parameter *CreateSubsysFromVolume) SetSrcVolumeName(srcVolumeName string) {
	parameter.srcVolumeName = utils.String(srcVolumeName)
	parameter.snapshotName = parameter.name
}

func (parameter *CreateSubsysFromVolume) EnableISCSI() {
	parameter.enableISCSI = utils.Bool(true)
}

func (parameter *CreateSubsysFromVolume) EnableNVMeoF() {
	parameter.enableNVMeoF = utils.Bool(true)
}

func (parameter *CreateSubsysFromVolume) SetSectorSize512() {
	parameter.sectorSize = utils.Int(512)
}

func (parameter *CreateSubsysFromVolume) SetSectorSize4096() {
	parameter.sectorSize = utils.Int(4096)
}
