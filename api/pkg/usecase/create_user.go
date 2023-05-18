package usecase

import (
	"context"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type createUser struct {
	repoUser repository.User
}

func NewCreateUser(ru repository.User) *createUser {
	return &createUser{repoUser: ru}
}

func (uc *createUser) Exec(ctx context.Context, name string) (*model.User, error) {
	u := model.NewUser(model.NewUserID(), name)

	err := uc.repoUser.Create(ctx, u)
	if err != nil {
		return nil, xerrors.Errorf("uc.repoUser.Create:  %v", err)
	}
	return u, nil
}
