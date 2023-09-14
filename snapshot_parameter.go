package prarameters

// File       : snapshot.go
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

type ListSnapshot struct {
	listParameter
	volumeName *string
}

func (parameter *ListSnapshot) GetVolumeName() string {
	return utils.StringValue(parameter.volumeName)
}

func (parameter *ListSnapshot) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *ListSnapshot) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}
	return fmt.Sprintf("snapshot/%s", *parameter.volumeName)
}

type RetrieveSnapshot struct {
	volumeName   *string
	snapshotName *string
}

func (parameter *RetrieveSnapshot) GetVolumeName() string {
	return utils.StringValue(parameter.volumeName)
}

func (parameter *RetrieveSnapshot) GetSnapshotName() string {
	return utils.StringValue(parameter.snapshotName)
}

func (parameter *RetrieveSnapshot) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *RetrieveSnapshot) SetSnapshotName(snapshotName string) {
	parameter.snapshotName = utils.String(snapshotName)
}

func (parameter *RetrieveSnapshot) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	if parameter.snapshotName == nil {
		panic("parameter snapshotName no set")
	}

	return fmt.Sprintf("snapshot/%s/%s", *parameter.volumeName, *parameter.snapshotName)
}

type DeleteSnapshot struct {
	volumeName   *string
	snapshotName *string
}

func (parameter *DeleteSnapshot) GetVolumeName() string {
	return utils.StringValue(parameter.volumeName)
}

func (parameter *DeleteSnapshot) GetSnapshotName() string {
	return utils.StringValue(parameter.volumeName)
}

func (parameter *DeleteSnapshot) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *DeleteSnapshot) SetSnapshotName(snapshotName string) {
	parameter.snapshotName = utils.String(snapshotName)
}

func (parameter *DeleteSnapshot) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	if parameter.snapshotName == nil {
		panic("parameter snapshotName no set")
	}

	return fmt.Sprintf("snapshot/%s/%s", *parameter.volumeName, *parameter.snapshotName)
}

type createDeleteSnapshot struct {
	volumeName   *string
	snapshotName *string
}

func (parameter *createDeleteSnapshot) MarshalJSON() ([]byte, error) {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	if parameter.snapshotName == nil {
		panic("parameter snapshotName no set")
	}

	return json.Marshal(map[string]*string{
		"volume_name":   parameter.volumeName,
		"snapshot_name": parameter.snapshotName,
	})
}

func (parameter *createDeleteSnapshot) GetVolumeName() string {
	return utils.StringValue(parameter.volumeName)
}

func (parameter *createDeleteSnapshot) GetSnapshotName() string {
	return utils.StringValue(parameter.volumeName)
}

func (parameter *createDeleteSnapshot) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *createDeleteSnapshot) SetSnapshotName(snapshotName string) {
	parameter.snapshotName = utils.String(snapshotName)
}

type CreateSnapshot = createDeleteSnapshot

type RevertSnapshot = createDeleteSnapshot
