package prarameters

// File       : common.go
// Path       : parameters
// Time       : CST 2023/4/26 14:15
// Group      : Taocloudx-FASS
// Author     : zhuc@taocloudx.com
// Description:

import (
	"strconv"

	"github.com/Vanishvei/fass-sdk-parameters/utils"
	"github.com/google/uuid"
)

type listParameter struct {
	Token    *string `form:"token"`
	PageSize *int    `form:"page_size"`
	PageNum  *int    `form:"page_num"`
}

func (parameter *listParameter) SetPageSize(pageSize int) {
	parameter.PageSize = utils.Int(pageSize)
}

func (parameter *listParameter) SetPageNum(pageNum int) {
	parameter.PageNum = utils.Int(pageNum)
}

func (parameter *listParameter) SetPageToken(token uuid.UUID) {
	parameter.Token = utils.String(token.String())
}

func (parameter *listParameter) GetPageSize() int {
	return utils.IntValue(parameter.PageSize)
}

func (parameter *listParameter) GetPageNum() int {
	return utils.IntValue(parameter.PageNum)
}

func (parameter *listParameter) GetQuery() map[string]*string {
	params := map[string]*string{}
	if parameter.PageNum != nil {
		params["page_num"] = utils.String(strconv.Itoa(*utils.Int(*parameter.PageNum)))
	}

	if parameter.PageSize != nil {
		params["page_size"] = utils.String(strconv.Itoa(*utils.Int(*parameter.PageSize)))
	}

	if parameter.Token != nil {
		params["token"] = parameter.Token
	}

	return params
}
