package prarameters

// File       : volume.go
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

type ListVolume = listParameter

type volumeParameter struct {
	volumeName *string
}

func (parameter *volumeParameter) GetVolumeName() string {
	return utils.StringValue(parameter.volumeName)
}

func (parameter *volumeParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

type ExpandVolume struct {
	volumeParameter
	capacity *string
}

func (parameter *ExpandVolume) MarshalJSON() ([]byte, error) {
	if parameter.capacity == nil {
		panic("parameter capacity no set")
	}
	return json.Marshal(map[string]string{"capacity": *parameter.capacity})
}

func (parameter *ExpandVolume) GetNewCapacity() string {
	return utils.StringValue(parameter.capacity)
}

func (parameter *ExpandVolume) SetNewCapacityGB(capacity int) {
	parameter.capacity = utils.String(fmt.Sprintf("%dg", capacity))
}

func (parameter *ExpandVolume) SetNewCapacityTB(capacity int) {
	parameter.capacity = utils.String(fmt.Sprintf("%dt", capacity))
}

func (parameter *ExpandVolume) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	return fmt.Sprintf("volume/%s/expand", *parameter.volumeName)
}

type RetrieveVolume struct {
	volumeParameter
}

func (parameter *RetrieveVolume) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	return fmt.Sprintf("volume/%s", *parameter.volumeName)
}

type DeleteVolume struct {
	volumeParameter
	isForce *bool
}

func (parameter *DeleteVolume) ForceDelete() {
	parameter.isForce = utils.Bool(true)
}

func (parameter *DeleteVolume) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}
	return fmt.Sprintf("volume/%s", *parameter.volumeName)
}

func (parameter *DeleteVolume) GetQuery() map[string]*string {
	if parameter.isForce == nil {
		return map[string]*string{"is_force": utils.String("false")}
	}
	return map[string]*string{"is_force": utils.String("true")}
}

type FlattenVolume struct {
	volumeParameter
}

func (parameter *FlattenVolume) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}
	return fmt.Sprintf("volume/%s/flatten", *parameter.volumeName)
}

type SetQosOfVolume struct {
	volumeParameter
	iops *int
	bps  *int
}

func (parameter *SetQosOfVolume) MarshalJSON() ([]byte, error) {
	_map := map[string]*int{}
	if parameter.iops == nil {
		panic("parameter iops no set")
	}
	if parameter.bps == nil {
		panic("parameter bps no set")
	}
	_map["iops"] = parameter.iops
	_map["bps"] = parameter.bps

	return json.Marshal(_map)
}

func (parameter *SetQosOfVolume) GetIOPS() int {
	return utils.IntValue(parameter.iops)
}

func (parameter *SetQosOfVolume) GetBPS() int {
	return utils.IntValue(parameter.bps)
}

func (parameter *SetQosOfVolume) SetIops(iops int) {
	parameter.iops = utils.Int(iops)
}

func (parameter *SetQosOfVolume) SetBpsMB(bps int) {
	parameter.bps = utils.Int(bps)
}

func (parameter *SetQosOfVolume) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	return fmt.Sprintf("volume/%s/set_qos", *parameter.volumeName)
}

type flattenVolume struct {
	taskId *string
}

func (parameter *flattenVolume) GetTaskId() string {
	return utils.StringValue(parameter.taskId)
}

func (parameter *flattenVolume) SetTaskId(taskId string) {
	parameter.taskId = utils.String(taskId)
}

type GetFlattenVolumeProgress struct {
	flattenVolume
}

func (parameter *GetFlattenVolumeProgress) GetPath() string {
	if parameter.taskId == nil {
		panic("parameter taskId no set")
	}
	return fmt.Sprintf("task/%s", *parameter.taskId)
}

type StopFlattenVolume struct {
	flattenVolume
}

func (parameter *StopFlattenVolume) GetPath() string {
	if parameter.taskId == nil {
		panic("parameter taskId no set")
	}
	return fmt.Sprintf("task/%s/stop", *parameter.taskId)
}
