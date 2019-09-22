package model

type User struct {
	Id       int    `xorm:"not null pk autoincr comment('用户id') INT(11)"`
	Username string `xorm:"not null comment('用户名') VARCHAR(32)"`
	Nickname string `xorm:"not null comment('昵称') VARCHAR(32)"`
	Age      int    `xorm:"not null comment('年龄') INT(11)"`
}
