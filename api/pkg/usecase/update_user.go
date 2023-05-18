package usecase

import (
	"context"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type updateUser struct {
	repoUser repository.User
}

func NewUpdateUser(ru repository.User) *updateUser {
	return &updateUser{repoUser: ru}
}

func (uc *updateUser) Exec(ctx context.Context, id, name string) error {
	if !model.IsValidUserID(id) {
		return xerrors.Errorf("!model.IsValidUserID: userID is invalid")
	}
	u := model.NewUser(id, name)

	err := uc.repoUser.Update(ctx, u)
	if err != nil {
		return xerrors.Errorf("uc.repoUser.Update: %v", err)
	}
	return nil
}
