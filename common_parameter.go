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
)

type listParameter struct {
	pageSize *int
	pageNum  *int
}

func (parameter *listParameter) SetPageSize(pageSize int) {
	parameter.pageSize = utils.Int(pageSize)
}

func (parameter *listParameter) SetPageNum(pageNum int) {
	parameter.pageNum = utils.Int(pageNum)
}

func (parameter *listParameter) GetPageSize() int {
	return utils.IntValue(parameter.pageSize)
}

func (parameter *listParameter) GetPageNum() int {
	return utils.IntValue(parameter.pageNum)
}

func (parameter *listParameter) GetQuery() map[string]*string {
	defaultPageNum := 1
	defaultPageSize := 20
	pageNum := strconv.Itoa(*utils.DefaultNumber(parameter.pageNum, utils.Int(defaultPageNum)))
	pageSize := strconv.Itoa(*utils.DefaultNumber(parameter.pageSize, utils.Int(defaultPageSize)))
	return map[string]*string{
		"page_num":  utils.String(pageNum),
		"page_size": utils.String(pageSize),
	}
}
