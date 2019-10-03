package myerr

// ParameterError 参数错误 自定义错误
type ParameterError struct {
	Message string
}

func (pe *ParameterError) Error() string {
	return pe.Message
}

// NewParameterError 创建参数错误 传入参数错误信息
func NewParameterError(message string) (parameterError *ParameterError) {
	return &ParameterError{message}
}
