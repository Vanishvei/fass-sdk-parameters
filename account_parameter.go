package prarameters

// File       : account.go
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

type ListAccount = listParameter

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

type RetrieveAccount = accountParameter

type DeleteAccount = accountParameter

type accountPasswdParameter struct {
	accountParameter
	password *string
}

func (parameter *accountPasswdParameter) MarshalJSON() ([]byte, error) {
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

func (parameter *accountPasswdParameter) GetPassword() string {
	return utils.StringValue(parameter.accountName)
}

func (parameter *accountPasswdParameter) SetPassword(password string) {
	parameter.password = utils.String(password)
}

type UpdateAccountPasswd struct {
	accountPasswdParameter
}

func (parameter *UpdateAccountPasswd) MarshalJSON() ([]byte, error) {
	return parameter.accountPasswdParameter.MarshalJSON()
}

func (parameter *UpdateAccountPasswd) GetPath() string {
	if parameter.accountName == nil {
		panic("parameter accountName no set")
	}
	return fmt.Sprintf("acl/account/%s/update_passwd", *parameter.accountName)
}

type CreateAccount struct {
	accountPasswdParameter
}

func (parameter *CreateAccount) MarshalJSON() ([]byte, error) {
	return parameter.accountPasswdParameter.MarshalJSON()
}

func (parameter *CreateAccount) GetPath() string {
	if parameter.accountName == nil {
		panic("parameter accountName no set")
	}
	return "acl/account"
}
