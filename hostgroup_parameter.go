package prarameters

// File       : hostgroup_parameter.go
// Path       :
// Time       : CST 2023/9/14 11:50
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"encoding/json"
	"fmt"

	"github.com/Vanishvei/fass-sdk-parameters/utils"
)

type ListHostGroup = listParameter

type hostGroupParameter struct {
	hostGroupName *string
}

func (parameter *hostGroupParameter) GetHostGroupName() string {
	return utils.StringValue(parameter.hostGroupName)
}

func (parameter *hostGroupParameter) SetHostGroupName(hostGroupName string) {
	parameter.hostGroupName = utils.String(hostGroupName)
}

type ListSubsysOfHostGroup struct {
	hostGroupParameter
}

func (parameter *ListSubsysOfHostGroup) GetPath() string {
	if parameter.hostGroupName == nil {
		panic("parameter hostGroupName no set")
	}
	return fmt.Sprintf("acl/host_group/%s/subsys", *parameter.hostGroupName)
}

type DeleteHostGroup struct {
	hostGroupParameter
}

func (parameter *DeleteHostGroup) GetPath() string {
	if parameter.hostGroupName == nil {
		panic("parameter HostGroupName no set")
	}
	return fmt.Sprintf("acl/host_group/%s", *parameter.hostGroupName)
}

type updateHostGroupParameter struct {
	hostGroupName *string
	iqn           *map[string]string
	nqn           *map[string]string
}

func (parameter *updateHostGroupParameter) MarshalJSON() ([]byte, error) {
	_map := map[string]interface{}{}
	_iqn := parameter.GetIQN()
	_nqn := parameter.GetNQN()
	if _iqn == nil && _nqn == nil {
		panic("parameter IQN and NQN cannot both be empty")
	}
	if _iqn != nil {
		_map["iqn"] = _iqn
	}
	if _nqn != nil {
		_map["nqn"] = _nqn
	}

	return json.Marshal(_map)
}

func (parameter *updateHostGroupParameter) GetHostGroupName() string {
	return utils.StringValue(parameter.hostGroupName)
}

func (parameter *updateHostGroupParameter) SetHostGroupName(hostGroupName string) {
	parameter.hostGroupName = utils.String(hostGroupName)
}

func (parameter *updateHostGroupParameter) GetIQN() map[string]string {
	return utils.MapValue(parameter.iqn)
}

func (parameter *updateHostGroupParameter) SetIQN(iqn map[string]string) {
	parameter.iqn = utils.Map(iqn)
}

func (parameter *updateHostGroupParameter) GetNQN() map[string]string {
	return utils.MapValue(parameter.nqn)
}

func (parameter *updateHostGroupParameter) SetNQN(nqn map[string]string) {
	parameter.nqn = utils.Map(nqn)
}

type AddHostToHostGroup struct {
	updateHostGroupParameter
}

func (parameter *AddHostToHostGroup) GetPath() string {
	if parameter.hostGroupName == nil {
		panic("parameter hostGroupName no set")
	}
	return fmt.Sprintf("acl/host_group/%s/add_hosts", *parameter.hostGroupName)
}

func (parameter *AddHostToHostGroup) MarshalJSON() ([]byte, error) {
	return parameter.updateHostGroupParameter.MarshalJSON()
}

type RemoveHostFromHostGroup struct {
	updateHostGroupParameter
}

func (parameter *RemoveHostFromHostGroup) GetPath() string {
	if parameter.hostGroupName == nil {
		panic("parameter hostGroupName no set")
	}
	return fmt.Sprintf("acl/host_group/%s/remove_hosts", *parameter.hostGroupName)
}

func (parameter *RemoveHostFromHostGroup) MarshalJSON() ([]byte, error) {
	return parameter.updateHostGroupParameter.MarshalJSON()
}
