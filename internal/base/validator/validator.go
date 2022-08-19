package validator

import (
	"bytes"
	"errors"
	"reflect"

	zhongwen "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
)

// MyValidator 自定义验证器
type MyValidator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

// GlobalValidator 全局验证器
var GlobalValidator MyValidator

func init() {
	zhs := zhongwen.New()
	uni := ut.New(zhs, zhs)
	trans, ok := uni.GetTranslator("zh")
	if !ok {
		panic(ok)
	}

	validate := validator.New()

	// 收集结构体中的comment标签，用于替换英文字段名称，这样返回错误就能展示中文字段名称了
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})

	// 注册中文翻译
	err := zh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	GlobalValidator.Validate = validate
	GlobalValidator.Trans = trans
}

// Check 验证器通用验证方法
func (m *MyValidator) Check(value interface{}) error {
	// 首先使用validator进行验证
	err := m.Validate.Struct(value)
	errBuf := bytes.Buffer{}
	if err != nil {
		var valErr validator.ValidationErrors
		if errors.As(err, &valErr) {
			// 将所有的参数错误进行翻译然后拼装成字符串返回
			for i := 0; i < len(valErr); i++ {
				errBuf.WriteString(valErr[i].Translate(m.Trans) + ", ")
			}
		} else {
			// 几乎不会出现，除非验证器本身异常无法转换，以防万一就判断一下好了
			return errors.New("validate check exception")
		}
	}

	// 如果它实现了CanCheck接口，就进行自定义验证
	if v, ok := value.(Checker); ok {
		errs := v.Check()
		for i := 0; i < len(errs); i++ {
			errBuf.WriteString(errs[i].Error() + ", ")
		}
	}

	if errBuf.Len() == 0 {
		return nil
	}

	// 删除掉最后一个空格和换行符
	errStr := errBuf.String()
	return errors.New(errStr[:len(errStr)-2])
}

// Checker 如果需要特殊校验，可以实现验证接口，或者通过自定义tag标签实现
type Checker interface {
	Check() []error
}
