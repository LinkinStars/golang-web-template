package pager

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPageModel(t *testing.T) {
	type args struct {
		pageNum      int
		pageSize     int
		totalRecords int64
		records      interface{}
	}
	tests := []struct {
		name string
		args args
		want *PageModel
	}{
		{
			"普通",
			args{
				pageNum:      0,
				pageSize:     0,
				totalRecords: -1,
				records:      make([]string, 0),
			},
			&PageModel{
				PageNum:      1,
				PageSize:     10,
				TotalPages:   0,
				TotalRecords: 0,
				Records:      make([]string, 0),
			},
		},
		{
			"异常",
			args{
				pageNum:      0,
				pageSize:     0,
				totalRecords: -1,
				records:      nil,
			},
			&PageModel{
				PageNum:      1,
				PageSize:     10,
				TotalPages:   0,
				TotalRecords: 0,
				Records:      make([]string, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if d := recover(); d != nil {
					fmt.Println(d)
				}
			}()
			if got := NewPageModel(tt.args.pageNum, tt.args.pageSize, tt.args.totalRecords, tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPageModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
