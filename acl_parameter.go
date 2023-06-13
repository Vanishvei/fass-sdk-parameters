package prarameters

// File       : acl.go
// Path       : parameters
// Time       : CST 2023/5/5 10:00
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"encoding/json"
	"fmt"

	"github.com/Vanishvei/fass-sdk-parameters/utils"
)

type ListAccountParameter = listParameter

type accountParameter struct {
	accountName *string
}

func (parameter *accountParameter) GetAccountName() string {
	return utils.StringValue(parameter.accountName)
}

func (parameter *accountParameter) SetAccountName(accountName string) {
	parameter.accountName = utils.String(accountName)
}

func (parameter *accountParameter) GetPath() string {
	if parameter.accountName == nil {
		panic("parameter accountName no set")
	}
	return fmt.Sprintf("acl/account/%s", *parameter.accountName)
}

type RetrieveAccountParameter = accountParameter

type DeleteAccountParameter = accountParameter

type CreateAccountParameter struct {
	accountName *string
	password    *string
}

func (parameter *CreateAccountParameter) GetAccountName() string {
	return utils.StringValue(parameter.accountName)
}

func (parameter *CreateAccountParameter) GetPassword() string {
	return utils.StringValue(parameter.password)
}

func (parameter *CreateAccountParameter) SetAccountName(accountName string) {
	parameter.accountName = utils.String(accountName)
}

func (parameter *CreateAccountParameter) SetPassword(password string) {
	parameter.password = utils.String(password)
}

func (parameter *CreateAccountParameter) MarshalJSON() ([]byte, error) {
	if parameter.accountName == nil {
		panic("parameter accountName no set")
	}

	if parameter.password == nil {
		panic("parameter password no set")
	}

	_map := map[string]interface{}{
		"account_name": *parameter.accountName,
		"password":     *parameter.password,
	}
	return json.Marshal(_map)
}

type ListGroupParameter = listParameter

type groupParameter struct {
	groupName *string
}

func (parameter *groupParameter) GetGroupName() string {
	return utils.StringValue(parameter.groupName)
}

func (parameter *groupParameter) SetGroupName(groupName string) {
	parameter.groupName = utils.String(groupName)
}

func (parameter *groupParameter) GetPath() string {
	if parameter.groupName == nil {
		panic("parameter groupName no set")
	}
	return fmt.Sprintf("acl/group/%s", *parameter.groupName)
}

type RetrieveGroupParameter = groupParameter

type DeleteGroupParameter = groupParameter

type AddQualifierToGroupParameter struct {
	groupName     *string
	hostName      *string
	qualifierList []*string
}

func (parameter *AddQualifierToGroupParameter) MarshalJSON() ([]byte, error) {
	_map := map[string]interface{}{
		"host_name":      *parameter.hostName,
		"qualifier_list": utils.StringSliceValue(parameter.qualifierList),
	}
	return json.Marshal(_map)
}

func (parameter *AddQualifierToGroupParameter) GetGroupName() string {
	return utils.StringValue(parameter.groupName)
}

func (parameter *AddQualifierToGroupParameter) GetHostName() string {
	return utils.StringValue(parameter.hostName)
}

func (parameter *AddQualifierToGroupParameter) GetQualifierList() []string {
	return utils.StringSliceValue(parameter.qualifierList)
}

func (parameter *AddQualifierToGroupParameter) SetGroupName(groupName string) {
	parameter.groupName = utils.String(groupName)
}

func (parameter *AddQualifierToGroupParameter) SetHostName(hostName string) {
	parameter.hostName = utils.String(hostName)
}

func (parameter *AddQualifierToGroupParameter) SetQualifierList(qualifierList []string) {
	parameter.qualifierList = utils.StringSlice(qualifierList)
}

func (parameter *AddQualifierToGroupParameter) GetPath() string {
	if parameter.groupName == nil {
		panic("parameter groupName no set")
	}
	return fmt.Sprintf("acl/group/%s/add_qualifier", *parameter.groupName)
}

type RemoveQualifierFromGroupParameter struct {
	groupName     *string
	hostName      *string
	qualifierList []*string
}

func (parameter *RemoveQualifierFromGroupParameter) MarshalJSON() ([]byte, error) {
	_map := map[string]interface{}{
		"host_name":      *parameter.hostName,
		"qualifier_list": utils.StringSliceValue(parameter.qualifierList),
	}
	return json.Marshal(_map)
}

func (parameter *RemoveQualifierFromGroupParameter) GetGroupName() string {
	return utils.StringValue(parameter.groupName)
}

func (parameter *RemoveQualifierFromGroupParameter) GetHostName() string {
	return utils.StringValue(parameter.hostName)
}

func (parameter *RemoveQualifierFromGroupParameter) GetQualifierList() []string {
	return utils.StringSliceValue(parameter.qualifierList)
}

func (parameter *RemoveQualifierFromGroupParameter) SetGroupName(groupName string) {
	parameter.groupName = utils.String(groupName)
}

func (parameter *RemoveQualifierFromGroupParameter) SetHostName(hostName string) {
	parameter.hostName = utils.String(hostName)
}

func (parameter *RemoveQualifierFromGroupParameter) SetQualifierList(qualifierList []string) {
	parameter.qualifierList = utils.StringSlice(qualifierList)
}

func (parameter *RemoveQualifierFromGroupParameter) GetPath() string {
	if parameter.groupName == nil {
		panic("parameter groupName no set")
	}
	return fmt.Sprintf("acl/group/%s/remove_qualifier", *parameter.groupName)
}
