package usecase

import (
	"context"
	"fmt"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
)

type GetUser struct {
	repoUser repository.User
}

func NewGetUser(ru repository.User) *GetUser {
	return &GetUser{
		repoUser: ru,
	}
}

func (uc *GetUser) Exec(ctx context.Context, id string) (*model.User, error) {
	if !model.IsValidUserID(id) {
		return nil, fmt.Errorf("usecase.GetUser: user id is invalid")
	}

	u, err := uc.repoUser.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("usecase.GetUser: %w", err)
	}
	return u, nil
}
