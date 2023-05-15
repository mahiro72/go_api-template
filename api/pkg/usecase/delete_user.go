package usecase

import (
	"context"
	"fmt"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
)

type DeleteUser struct {
	repoUser repository.User
}

func NewDeleteUser(ru repository.User) *DeleteUser {
	return &DeleteUser{
		repoUser: ru,
	}
}

func (uc *DeleteUser) Exec(ctx context.Context, id string) error {
	if !model.IsValidUserID(id) {
		return fmt.Errorf("usecase.DeleteUser: user id is invalid")
	}

	err := uc.repoUser.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("usecase.DeleteUser: %w", err)
	}
	return nil
}
