package pager

import (
	"reflect"
)

// 分页模型
type PageModel struct {
	PageNum      int         `json:"page_num"`
	PageSize     int         `json:"page_size"`
	TotalPages   int64       `json:"total_pages"`
	TotalRecords int64       `json:"total_records"`
	Records      interface{} `json:"records"`
}

// 构造分页对象
// pageNum: 当前页码，从1开始
// pageSize: 每页条数
// totalRecords: 总条数
// records: 分页数据，必须为slice，不能为nil或者其他对象，可以是空的slice
func NewPageModel(pageNum, pageSize int, totalRecords int64, records interface{}) *PageModel {
	sliceValue := reflect.Indirect(reflect.ValueOf(records))
	if sliceValue.Kind() != reflect.Slice {
		panic("分页异常，需要传入一个slice类型")
	}
	if pageSize <= 0 {
		panic("分页异常，每页条数必须为正数")
	}

	if pageNum <= 0 {
		pageNum = 1
	}
	if totalRecords < 0 {
		totalRecords = 0
	}

	totalPages := calcTotalPages(int64(pageSize), totalRecords)

	return &PageModel{
		PageNum:      pageNum,
		PageSize:     pageSize,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
		Records:      records,
	}
}

// 计算页码逻辑
func calcTotalPages(pageSize, totalRecords int64) (totalPages int64) {
	if totalRecords == 0 {
		return 0
	}

	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return totalPages
}
