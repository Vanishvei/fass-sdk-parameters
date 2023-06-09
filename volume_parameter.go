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

type ListVolumeParameter = listParameter

type ExpandVolumeParameter struct {
	volumeName *string
	capacity   *string
}

func (parameter *ExpandVolumeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"capacity": *parameter.capacity})
}

func (parameter *ExpandVolumeParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *ExpandVolumeParameter) SetCapacityGB(capacity int) {
	parameter.capacity = utils.String(fmt.Sprintf("%dg", capacity))
}

func (parameter *ExpandVolumeParameter) SetCapacityTB(capacity int) {
	parameter.capacity = utils.String(fmt.Sprintf("%dt", capacity))
}

func (parameter *ExpandVolumeParameter) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	return fmt.Sprintf("volume/%s", *parameter.volumeName)
}

type RetrieveVolumeParameter struct {
	volumeName *string
}

func (parameter *RetrieveVolumeParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *RetrieveVolumeParameter) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	return fmt.Sprintf("volume/%s", *parameter.volumeName)
}

type DeleteVolumeParameter struct {
	volumeName *string
	isForce    *bool
}

func (parameter *DeleteVolumeParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *DeleteVolumeParameter) ForceDelete() {
	parameter.isForce = utils.Bool(true)
}

func (parameter *DeleteVolumeParameter) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}
	return fmt.Sprintf("volume/%s", *parameter.volumeName)
}

func (parameter *DeleteVolumeParameter) GetQuery() map[string]*string {
	if parameter.isForce == nil {
		return map[string]*string{"is_force": utils.String("false")}
	}
	return map[string]*string{"is_force": utils.String("true")}
}

type FlattenVolumeParameter struct {
	volumeName *string
}

func (parameter *FlattenVolumeParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *FlattenVolumeParameter) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	return fmt.Sprintf("volume/%s/flatten", *parameter.volumeName)
}

type SetQosOfVolumeParameter struct {
	volumeName *string
	iops       *int
	bps        *int
}

func (parameter *SetQosOfVolumeParameter) MarshalJSON() ([]byte, error) {
	_map := map[string]*int{}
	if parameter.iops != nil {
		panic("parameter iops no set")
	}
	if parameter.bps != nil {
		panic("parameter bps no set")
	}
	_map["iops"] = parameter.iops
	_map["bps"] = parameter.bps

	return json.Marshal(_map)
}

func (parameter *SetQosOfVolumeParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *SetQosOfVolumeParameter) SetIops(iops int) {
	parameter.iops = utils.Int(iops)
}

func (parameter *SetQosOfVolumeParameter) SetBpsMB(bps int) {
	parameter.bps = utils.Int(bps)
}

func (parameter *SetQosOfVolumeParameter) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	return fmt.Sprintf("volume/%s", *parameter.volumeName)
}

type flattenVolume struct {
	taskId *string
}

type GetFlattenVolumeProgress flattenVolume

func (parameter *GetFlattenVolumeProgress) SetTaskId(taskId string) {
	parameter.taskId = utils.String(taskId)
}

func (parameter *GetFlattenVolumeProgress) GetPath() string {
	if parameter.taskId == nil {
		panic("parameter taskId no set")
	}
	return fmt.Sprintf("task/%s", *parameter.taskId)
}

type StopFlattenVolumeParameter flattenVolume

func (parameter *StopFlattenVolumeParameter) SetTaskId(taskId string) {
	parameter.taskId = utils.String(taskId)
}

func (parameter *StopFlattenVolumeParameter) GetPath() string {
	if parameter.taskId == nil {
		panic("parameter taskId no set")
	}
	return fmt.Sprintf("task/%s/stop", *parameter.taskId)
}
