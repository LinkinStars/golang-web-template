package pager

import (
	"errors"
	"reflect"
)

// 分页模型
type PageModel struct {
	Page         int         `json:"page"`
	PageSize     int         `json:"page_size"`
	TotalPages   int64       `json:"total_pages"`
	TotalRecords int64       `json:"total_records"`
	Records      interface{} `json:"records"`
}

// 构造分页对象
// page: 当前页码
// pageSize: 每页条数
// totalRecords: 总条数
// records: 分页数据，必须为slice，不能为nil或者其他对象，可以是空的slice
func NewPageModel(page, pageSize int, totalRecords int64, records interface{}) *PageModel {
	sliceValue := reflect.Indirect(reflect.ValueOf(records))
	if sliceValue.Kind() != reflect.Slice {
		panic("分页异常，需要传入一个slice类型")
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if totalRecords < 0 {
		totalRecords = 0
	}

	totalPages, _ := calcTotalPages(int64(pageSize), totalRecords)

	return &PageModel{
		Page:         page,
		PageSize:     pageSize,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
		Records:      records,
	}
}

// 计算页码逻辑
func calcTotalPages(pageSize, totalRecords int64) (totalPages int64, err error) {
	if pageSize <= 0 {
		return 0, errors.New("page size should be positive integer")
	}

	if totalRecords < 0 {
		return 0, errors.New("total records should not be negative integer")
	}

	if totalRecords == 0 {
		return 0, nil
	}

	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return totalPages, nil
}
