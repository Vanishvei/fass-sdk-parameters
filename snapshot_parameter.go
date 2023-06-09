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

type ListSnapshotParameter struct {
	volumeName *string
	pageSize   *int
	pageNum    *int
}

func (parameter *ListSnapshotParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *ListSnapshotParameter) SetPageSize(pageSize int) {
	parameter.pageSize = utils.Int(pageSize)
}

func (parameter *ListSnapshotParameter) SetPageNum(pageNum int) {
	parameter.pageNum = utils.Int(pageNum)
}

func (parameter *ListSnapshotParameter) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}
	return fmt.Sprintf("/snapshot/%s", *parameter.volumeName)
}

func (parameter *ListSnapshotParameter) GetQuery() map[string]*int {
	_map := map[string]*int{}
	if parameter.pageSize != nil {
		_map["page_size"] = parameter.pageSize
	}
	if parameter.pageNum != nil {
		_map["page_num"] = parameter.pageNum
	}
	return _map
}

type RetrieveSnapshotParameter struct {
	volumeName   *string
	snapshotName *string
}

func (parameter *RetrieveSnapshotParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *RetrieveSnapshotParameter) SetSnapshotName(snapshotName string) {
	parameter.snapshotName = utils.String(snapshotName)
}

func (parameter *RetrieveSnapshotParameter) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	if parameter.snapshotName == nil {
		panic("parameter snapshotName no set")
	}

	return fmt.Sprintf("snapshot/%s/%s", *parameter.volumeName, *parameter.snapshotName)
}

type DeleteSnapshotParameter struct {
	volumeName   *string
	snapshotName *string
}

func (parameter *DeleteSnapshotParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *DeleteSnapshotParameter) SetSnapshotName(snapshotName string) {
	parameter.snapshotName = utils.String(snapshotName)
}

func (parameter *DeleteSnapshotParameter) GetPath() string {
	if parameter.volumeName == nil {
		panic("parameter volumeName no set")
	}

	if parameter.snapshotName == nil {
		panic("parameter snapshotName no set")
	}

	return fmt.Sprintf("snapshot/%s/%s", *parameter.volumeName, *parameter.snapshotName)
}

type createDeleteSnapshotParameter struct {
	volumeName   *string
	snapshotName *string
}

func (parameter *createDeleteSnapshotParameter) MarshalJSON() ([]byte, error) {
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

func (parameter *createDeleteSnapshotParameter) SetVolumeName(volumeName string) {
	parameter.volumeName = utils.String(volumeName)
}

func (parameter *createDeleteSnapshotParameter) SetSnapshotName(snapshotName string) {
	parameter.snapshotName = utils.String(snapshotName)
}

type CreateSnapshotParameter = createDeleteSnapshotParameter

type RevertSnapshotParameter = createDeleteSnapshotParameter
