package mistake

import (
	"base/logger"
)

// 请求参数异常
type ReqErr struct {
	Message string
}

func (e *ReqErr) Error() string {
	return e.Message
}

func NewReqErr(message string) *ReqErr {
	return &ReqErr{Message: message}
}

// 业务服务异常
type ServiceErr struct {
	Message string
	Err     error
	Stack   string
}

func (e *ServiceErr) Error() string {
	return e.Message
}

func NewServiceErr(err error, message string) *ServiceErr {
	return &ServiceErr{
		Message: message,
		Err:     err,
		Stack:   logger.LogStack(2, 5),
	}
}

// 数据层异常
type DaoErr struct {
	Err   error
	Stack string
}

func (e *DaoErr) Error() string {
	return e.Err.Error()
}

func NewDaoErr(err error) *DaoErr {
	return &DaoErr{
		Err:   err,
		Stack: logger.LogStack(2, 5),
	}
}
