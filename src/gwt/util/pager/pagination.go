package pager

import (
	"errors"
	"reflect"

	"github.com/go-xorm/xorm"
)

// PageModel 分页模型
type PageModel struct {
	PageNum      int         `json:"page_num"`
	PageSize     int         `json:"page_size"`
	TotalPages   int64       `json:"total_pages"`
	TotalRecords int64       `json:"total_records"`
	Records      interface{} `json:"records"`
}

// Help xorm 分页
// pageNum: 当前页码，从1开始
// pageSize: 每页条数
// rowsSlicePtr: 要求传入一个slice的指针类型 &[]model.User，否则出现错误
// rowElement: slice里面的单个元素，会将其中的非0值元素作为查询条件
// session: 为带有查询条件的session 可以携带一些特殊的查询条件 rowElement 无法做到的，如：like 等
// 默认分页大小20，如果查询总数为0则不会继续往后查询
// 会赋值传入的slice为查询结果，并返回查询总数total，或执行中出现的error错误
func Help(pageNum, pageSize int, rowsSlicePtr interface{}, rowElement interface{}, session *xorm.Session) (total int64, err error) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	// 保证传入的是slice类型
	sliceValue := reflect.Indirect(reflect.ValueOf(rowsSlicePtr))
	if sliceValue.Kind() != reflect.Slice {
		return 0, errors.New("需要传入一个slice类型")
	}

	// clone一个session为了去执行第二个sql语句
	tempSession := session.Clone()

	// 查询总数，如果总数为0，就不继续查询了
	total, err = session.Count(rowElement)
	if err != nil || total == 0 {
		return
	}

	// 查询具体数据
	startNum := (pageNum - 1) * pageSize
	err = tempSession.Limit(pageSize, startNum).Find(rowsSlicePtr, rowElement)
	return
}

// NewPageModel 构造分页对象
// pageNum: 当前页码，从1开始
// pageSize: 每页条数，默认条数20
// totalRecords: 总条数
// records: 分页数据，必须为slice，不能为nil或者其他对象，可以是空的slice
func NewPageModel(pageNum, pageSize int, totalRecords int64, records interface{}) *PageModel {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 20
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

// calcTotalPages 计算页码逻辑
func calcTotalPages(pageSize, totalRecords int64) (totalPages int64) {
	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return totalPages
}
