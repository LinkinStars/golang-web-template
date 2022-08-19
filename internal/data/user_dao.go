package data

import (
	"context"
	"time"

	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/go-scaffold/mistake"
	"github.com/LinkinStars/golang-web-template/internal/biz"
	"github.com/LinkinStars/golang-web-template/src/gwt/util/pager"
	"github.com/jinzhu/copier"
)

// User 用户
type User struct {
	ID        int       `xorm:"not null pk autoincr comment('用户ID') INT(11) id"`
	CreatedAt time.Time `xorm:"created comment('创建时间') TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated comment('修改时间') TIMESTAMP updated_at"`
	Username  string    `xorm:"not null comment('用户名') VARCHAR(32) username"`
	Nickname  string    `xorm:"not null comment('昵称') VARCHAR(16) nickname"`
	Age       int       `xorm:"not null comment('年龄') INT(11) age"`
}

// TableName 用户 表名
func (User) TableName() string {
	return "user"
}

type userRepo struct {
	data *Data
	log  logger.Logger
}

func (u userRepo) AddUser(ctx context.Context, userInfo *biz.UserDTO) (err error) {
	user := &User{}
	_ = copier.Copy(user, userInfo)
	_, err = u.data.db.Insert(user)
	if err != nil {
		return mistake.InternalServer("add user failed", "internal error").WithError(err).WithStack()
	}
	return
}

func (u userRepo) RemoveUser(ctx context.Context, id int) (err error) {
	_, err = u.data.db.ID(id).Delete(&User{ID: id})
	if err != nil {
		return mistake.InternalServer("remove user failed", "internal error").WithError(err).WithStack()
	}
	return
}

func (u userRepo) UpdateUser(ctx context.Context, userInfo *biz.UserDTO) (err error) {
	user := &User{}
	_ = copier.Copy(user, userInfo)
	_, err = u.data.db.ID(user.ID).Update(user)
	if err != nil {
		return mistake.InternalServer("update user failed", "internal error").WithError(err).WithStack()
	}
	return
}

func (u userRepo) GetUser(ctx context.Context, id int) (userInfo *biz.UserDTO, exist bool, err error) {
	user := &User{}
	exist, err = u.data.db.ID(id).Get(user)
	if err != nil {
		return nil, false, mistake.InternalServer("get user failed", "internal error").WithError(err).WithStack()
	}
	if !exist {
		return nil, false, nil
	}
	userInfo = &biz.UserDTO{}
	_ = copier.Copy(userInfo, user)
	return
}

func (u userRepo) GetUsers(ctx context.Context, userInfo *biz.UserDTO) (userList []*biz.UserDTO, err error) {
	user := &User{}
	_ = copier.Copy(user, userInfo)
	var users []*User
	err = u.data.db.Find(&users, user)
	if err != nil {
		return nil, mistake.InternalServer("get users failed", "internal error").WithError(err).WithStack()
	}
	userList = make([]*biz.UserDTO, 0)
	_ = copier.Copy(&userList, users)
	return
}

func (u userRepo) GetUsersWithPage(ctx context.Context, page, pageSize int, userInfo *biz.UserDTO) (
	userList []*biz.UserDTO, total int64, err error) {
	user := &User{}
	_ = copier.Copy(user, userInfo)
	var users []*User
	total, err = pager.Help(page, pageSize, &users, user, u.data.db.NewSession())
	if err != nil {
		return nil, 0, mistake.InternalServer("get users with page failed", "internal error").WithError(err).WithStack()
	}

	userList = make([]*biz.UserDTO, 0)
	_ = copier.Copy(&userList, users)
	return
}

// NewUserRepo .
func NewUserRepo(data *Data, log logger.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log,
	}
}
