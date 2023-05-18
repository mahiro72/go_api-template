package usecase

import (
	"context"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type deleteUser struct {
	repoUser repository.User
}

func NewDeleteUser(ru repository.User) *deleteUser {
	return &deleteUser{repoUser: ru}
}

func (uc *deleteUser) Exec(ctx context.Context, id string) error {
	if !model.IsValidUserID(id) {
		return xerrors.Errorf("!model.IsValidUserID: userID is invalid")
	}

	err := uc.repoUser.Delete(ctx, id)
	if err != nil {
		return xerrors.Errorf("uc.repoUser.Delete: %v", err)
	}
	return nil
}
