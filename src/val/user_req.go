package val

import (
	"errors"
	"strings"
)

type AddUserReq struct {
	Username string `validate:"required,gt=6,lt=16" comment:"用户名"`
	Nickname string `validate:"required,gt=6,lt=16" comment:"昵称"`
	Age      int    `validate:"required,min=5,max=99" comment:"年龄"`
}

func (u *AddUserReq) Check() error {
	if strings.Contains(u.Username, "@") {
		return errors.New("用户名不能包含@特殊字符")
	}
	return nil
}

type DeleteUserReq struct {
	Id int `validate:"required,min=1" comment:"Id"`
}

type UpdateUserReq struct {
	Id       int    `validate:"required,min=1" comment:"Id"`
	Username string `validate:"required,gt=6,lt=16" comment:"用户名"`
	Nickname string `validate:"required,gt=6,lt=16" comment:"昵称"`
	Age      int    `validate:"required,min=5,max=99" comment:"年龄"`
}

type GetUserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}
