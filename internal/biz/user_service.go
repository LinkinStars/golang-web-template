package biz

import (
	"context"
	"time"

	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/LinkinStars/go-scaffold/mistake"
	"github.com/LinkinStars/golang-web-template/internal/base/pager"
	"github.com/LinkinStars/golang-web-template/internal/val"
	"github.com/jinzhu/copier"
)

type UserRepo interface {
	AddUser(ctx context.Context, userInfo *UserDTO) error
	RemoveUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, userInfo *UserDTO) error
	GetUser(ctx context.Context, id int) (*UserDTO, bool, error)
	GetUsers(ctx context.Context, userInfo *UserDTO) ([]*UserDTO, error)
	GetUsersWithPage(ctx context.Context, page, pageSize int, userInfo *UserDTO) (
		users []*UserDTO, total int64, err error)
}

type UserUseCase struct {
	repo UserRepo
	log  logger.Logger
}

func NewUserUseCase(repo UserRepo, log logger.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log}
}

type UserDTO struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Nickname  string
	Age       int
}

// AddUser 新增用户
func (u *UserUseCase) AddUser(ctx context.Context, req *val.AddUserReq) (err error) {
	user := &UserDTO{}
	_ = copier.Copy(user, req)
	err = u.repo.AddUser(ctx, user)
	if err != nil {
		return err
	}
	return
}

// RemoveUser 删除用户
func (u *UserUseCase) RemoveUser(ctx context.Context, id int) (err error) {
	err = u.repo.RemoveUser(ctx, id)
	if err != nil {
		return err
	}
	return
}

// UpdateUser 修改用户
func (u *UserUseCase) UpdateUser(ctx context.Context, req *val.UpdateUserReq) (err error) {
	user := &UserDTO{}
	_ = copier.Copy(user, req)
	err = u.repo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return
}

// GetUser 查询用户 单个
func (u *UserUseCase) GetUser(ctx context.Context, id int) (req *val.GetUserResp, err error) {
	user, exist, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, mistake.BadRequest("user not exist", "user not exist").WithStack()
	}

	req = &val.GetUserResp{}
	_ = copier.Copy(req, user)
	return req, nil
}

// GetUsers 查询用户列表 全部
func (u *UserUseCase) GetUsers(ctx context.Context, req *val.GetUsersReq) (usersResp []*val.GetUserResp, err error) {
	user := &UserDTO{}
	_ = copier.Copy(user, req)

	users, err := u.repo.GetUsers(ctx, user)
	if err != nil {
		return nil, err
	}

	usersResp = make([]*val.GetUserResp, 0)
	_ = copier.Copy(&usersResp, users)
	return
}

// GetUsersWithPage 查询用户列表 分页
func (u *UserUseCase) GetUsersWithPage(ctx context.Context, req *val.GetUsersWithPageReq) (pageModel *pager.PageModel, err error) {
	user := &UserDTO{}
	_ = copier.Copy(user, req)

	page := req.Page
	pageSize := req.PageSize

	users, total, err := u.repo.GetUsersWithPage(ctx, page, pageSize, user)
	if err != nil {
		return nil, err
	}

	usersResp := make([]*val.GetUserResp, 0)
	_ = copier.Copy(&usersResp, users)

	return pager.NewPageModel(page, pageSize, total, usersResp), nil
}
