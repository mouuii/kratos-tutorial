package data

import (
	"context"
	"devops-agent/internal/biz"
	"devops-agent/internal/data/ent/user"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	fmt.Println(r.data.conf.Cit)
	_, err := r.data.db.User.Create().SetName(u.Name).SetAge(int(u.Age)).Save(ctx)
	return u, err
}

func (r *userRepo) Update(ctx context.Context, u *biz.User) (*biz.User, error) {
	return u, nil
}

func (r *userRepo) FindByID(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListByUserName(ctx context.Context, s string) (bool, error) {
	return r.data.db.User.Query().Where(user.NameEQ(s)).Exist(ctx)
}

func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}
