package pager

import (
	"errors"
	"reflect"

	"github.com/go-xorm/xorm"
)

// xorm 分页
// page: 当前页码
// pageSize: 每页条数
// rowsSlicePtr: 要求传入一个slice的指针类型 &[]model.User，否则出现错误
// session: 为带有查询条件的session
// 如果查询总数为0则不会继续往后查询
// 会赋值传入的slice为查询结果，并返回查询总数total，或执行中出现的error错误
func Help(page, pageSize int, rowsSlicePtr interface{}, session *xorm.Session) (total int64, err error) {
	// 保证传入的是slice类型
	sliceValue := reflect.Indirect(reflect.ValueOf(rowsSlicePtr))
	if sliceValue.Kind() != reflect.Slice {
		return 0, errors.New("需要传入一个slice类型")
	}

	// 获取slice内部元素类型，如果内部还是指针类型就拿取指针类型的值
	sliceElementType := sliceValue.Type().Elem()
	if sliceElementType.Kind() == reflect.Ptr {
		sliceElementType = sliceElementType.Elem()
	}

	// clone一个session为了去执行第二个sql语句
	tempSession := session.Clone()

	// 查询总数，如果总数为0，就不继续查询了
	total, err = session.Count(reflect.New(sliceElementType).Interface())
	if err != nil || total == 0 {
		return
	}

	// 查询具体数据
	startNum := (page - 1) * pageSize
	err = tempSession.Limit(pageSize, startNum).Find(rowsSlicePtr)
	return
}
