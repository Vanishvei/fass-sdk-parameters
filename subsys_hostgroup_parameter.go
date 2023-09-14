package prarameters

// File       : subsys_hostgroup_parameter.go
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

type RetrieveSubsysAuth struct {
	subsysParameter
}

func (parameter *RetrieveSubsysAuth) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/host_group", *parameter.subsysName)
}

type SubsysBindHostGroup struct {
	subsysParameter
	hostGroupName *string
}

func (parameter *SubsysBindHostGroup) SetHostGroupName(groupName string) {
	parameter.hostGroupName = utils.String(groupName)
}

func (parameter *SubsysBindHostGroup) GetHostGroupName() string {
	return utils.StringValue(parameter.hostGroupName)
}

func (parameter *SubsysBindHostGroup) MarshalJSON() ([]byte, error) {
	if parameter.hostGroupName == nil {
		panic("parameter group_name no set")
	}
	return json.Marshal(map[string]string{"group_name": *parameter.hostGroupName})
}

func (parameter *SubsysBindHostGroup) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/host_group/bind", *parameter.subsysName)
}

type SubsysUnbindAuth struct {
	subsysParameter
}

func (parameter *SubsysUnbindAuth) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/auth/remove", *parameter.subsysName)
}
