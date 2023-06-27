package biz

import (
	"context"

	v1 "devops-agent/api/agent/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound     = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
	ErrUserAlreadyExist = errors.New(500, v1.ErrorReason_USER_ALREADY_EXIST.String(), "user already exist")
)

// User is a User model.
type User struct {
	Name string
	Age  int32
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListByUserName(context.Context, string) (bool, error)
	ListAll(context.Context) ([]*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Register creates a User, and returns the new User.
func (uc *UserUsecase) Register(ctx context.Context, u *User) (*User, error) {
	// 查询数据库是否有相同的用户名，如果有返回错误
	uc.log.WithContext(ctx).Infof("Register: %v", u.Name)
	isExist, err := uc.repo.ListByUserName(ctx, u.Name)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, ErrUserAlreadyExist
	}
	return uc.repo.Save(ctx, u)
}
