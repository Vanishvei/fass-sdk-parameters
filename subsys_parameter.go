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

type ListSubsysParameter = listParameter

type RetrieveSubsysParameter struct {
	subsysName *string
}

func (parameter *RetrieveSubsysParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
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

func (parameter ExportSubsysParameter) GetPath() string {
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

func (parameter UnexportSubsysParameter) GetPath() string {
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

func (parameter CreateSubsysParameter) MarshalJSON() ([]byte, error) {
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

	if parameter.srcVolumeName != nil && parameter.snapshotName == nil {
		panic("parameter srcSnapshotName no set")
	}

	if parameter.snapshotName != nil && parameter.srcVolumeName == nil {
		panic("parameter srcVolumeName no set")
	}

	_map := map[string]interface{}{
		"name":          *parameter.name,
		"pool_name":     *parameter.poolName,
		"protocol_type": *parameter.protocolType,
	}

	if parameter.bps != nil {
		_map["bps"] = *parameter.bps
	}
	if parameter.capacity != nil {
		_map["capacity"] = *parameter.capacity
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

	if parameter.snapshotName != nil {
		_map["snapshot_name"] = *parameter.snapshotName
	}

	if parameter.srcVolumeName != nil {
		_map["src_volume_name"] = *parameter.srcVolumeName
	}

	if parameter.volumeName != nil {
		_map["volume_name"] = *parameter.volumeName
	}

	return json.Marshal(_map)
}

func (parameter *CreateSubsysParameter) SetName(subsysName string) {
	parameter.name = utils.String(subsysName)
}

func (parameter *CreateSubsysParameter) SetBpsMB(bps int) {
	parameter.bps = utils.Int(bps)
}

func (parameter *CreateSubsysParameter) SetIops(iops int) {
	parameter.iops = utils.Int(iops)
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

func (parameter *CreateSubsysParameter) InheritQos() {
	parameter.inheritQos = utils.Bool(true)
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

func (parameter *CreateSubsysParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *CreateSubsysParameter) SetSrcVolumeName(srcVolumeName string) {
	parameter.srcVolumeName = utils.String(srcVolumeName)
}

func (parameter *CreateSubsysParameter) SetSrcSnapshotName(srcSnapshotName string) {
	parameter.snapshotName = utils.String(srcSnapshotName)
}

type RetrieveSubsysAuthParameter struct {
	subsysName *string
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

func (parameter SetSubsysAuthParameter) MarshalJSON() ([]byte, error) {
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

func (parameter SetSubsysChapParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"account_name": *parameter.accountName})
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

func (parameter *RemoveSubsysChapParameter) SetSubsysName(subsysName string) {
	parameter.subsysName = utils.String(subsysName)
}

func (parameter *RemoveSubsysChapParameter) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/chap/remove", *parameter.subsysName)
}
