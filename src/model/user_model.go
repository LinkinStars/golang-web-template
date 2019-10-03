package model

// User 用户
type User struct {
	ID       int    `xorm:"not null pk autoincr comment('用户id') INT(11) id"`
	Username string `xorm:"not null comment('用户名') VARCHAR(32) username"`
	Nickname string `xorm:"comment('昵称') VARCHAR(32) nickname"`
	Age      int    `xorm:"not null comment('年龄') INT(11) age"`
}
