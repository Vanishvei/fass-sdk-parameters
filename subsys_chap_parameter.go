package prarameters

// File       : subsys_chap_parameter.go
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

type RetrieveSubsysChap struct {
	subsysParameter
}

func (parameter *RetrieveSubsysChap) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/chap", *parameter.subsysName)
}

type SetSubsysChap struct {
	subsysParameter
	accountName *string
}

func (parameter *SetSubsysChap) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"account_name": *parameter.accountName})
}

func (parameter *SetSubsysChap) GetAccountName() string {
	return utils.StringValue(parameter.accountName)
}

func (parameter *SetSubsysChap) SetAccountName(subsysName string) {
	parameter.accountName = utils.String(subsysName)
}

func (parameter *SetSubsysChap) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/chap/set", *parameter.subsysName)
}

type RemoveSubsysChap struct {
	subsysParameter
}

func (parameter *RemoveSubsysChap) GetPath() string {
	if parameter.subsysName == nil {
		panic("parameter subsysName no set")
	}
	return fmt.Sprintf("subsys/%s/chap/remove", *parameter.subsysName)
}
