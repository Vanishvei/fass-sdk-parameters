package prarameters

// File       : subsys_vlan_parameter.go
// Path       :
// Time       : CST 2023/9/14 15:04
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"encoding/json"
	"fmt"

	"github.com/Vanishvei/fass-sdk-parameters/utils"
)

type vlan struct {
	subsysParameter
}

func (parameter *vlan) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/vlan", *parameter.subsysName)
}

type RetrieveSubsysVLAN = vlan

type DeleteSubsysVLAN = vlan

type subsysVLAN struct {
	subsysParameter
	vlanList []*string
}

func (parameter *subsysVLAN) MarshalJSON() ([]byte, error) {
	if parameter.vlanList == nil {
		panic("parameter vlan_list no set")
	}
	if len(parameter.vlanList) == 0 {
		panic("parameter vlan_list no set")
	}
	return json.Marshal(map[string]any{"vlan_list": utils.StringSliceValue(parameter.vlanList)})
}

func (parameter *subsysVLAN) GetVLANList() []string {
	return utils.StringSliceValue(parameter.vlanList)
}

func (parameter *subsysVLAN) SetVLANList(vlanList []string) {
	parameter.vlanList = utils.StringSlice(vlanList)
}

type SubsysAddVLAN struct {
	subsysVLAN
}

func (parameter *SubsysAddVLAN) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/vlan/add", *parameter.subsysName)
}

type SubsysRemoveVLAN struct {
	subsysVLAN
}

func (parameter *SubsysRemoveVLAN) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/vlan/remove", *parameter.subsysName)
}
