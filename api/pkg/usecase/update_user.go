package usecase

import (
	"context"
	"fmt"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
)

type UpdateUser struct {
	repoUser repository.User
}

func NewUpdateUser(ru repository.User) *UpdateUser {
	return &UpdateUser{
		repoUser: ru,
	}
}

func (uc *UpdateUser) Exec(ctx context.Context, id, name string) error {
	if !model.IsValidUserID(id) {
		return fmt.Errorf("usecase.UpdateUser: user id is invalid")
	}

	err := uc.repoUser.Update(ctx, id, name)
	if err != nil {
		return fmt.Errorf("usecase.UpdateUser: %w", err)
	}
	return nil
}
