package usecase

import (
	"context"
	"fmt"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
)

type CreateUser struct {
	repoUser repository.User
}

func NewCreateUser(ru repository.User) *CreateUser {
	return &CreateUser{
		repoUser: ru,
	}
}

func (uc *CreateUser) Exec(ctx context.Context, name string) (*model.User, error) {
	id := model.NewUserID()

	err := uc.repoUser.Create(ctx, id, name)
	if err != nil {
		return nil, fmt.Errorf("usecase.CreateUser: %w", err)
	}
	return model.NewUser(id, name), nil
}
