package prarameters

// File       : pool.go
// Path       : parameters
// Time       : CST 2023/4/25 14:59
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"fmt"

	"github.com/Vanishvei/fass-sdk-parameters/utils"
)

type RetrievePoolParameter struct {
	poolName *string
}

func (parameter *RetrievePoolParameter) GetPoolName() string {
	return utils.StringValue(parameter.poolName)
}

func (parameter *RetrievePoolParameter) SetPoolName(poolName string) {
	parameter.poolName = utils.String(poolName)
}

func (parameter *RetrievePoolParameter) GetPath() string {
	if parameter.poolName == nil {
		panic("parameter poolName not set")
	}
	return fmt.Sprintf("pool/%s", *parameter.poolName)
}
