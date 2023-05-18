package usecase

import (
	"context"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type getUser struct {
	repoUser repository.User
}

func NewGetUser(ru repository.User) *getUser {
	return &getUser{repoUser: ru}
}

func (uc *getUser) Exec(ctx context.Context, id string) (*model.User, error) {
	if !model.IsValidUserID(id) {
		return nil, xerrors.Errorf("!model.IsValidUserID: userID is invalid")
	}
	u, err := uc.repoUser.Get(ctx, id)
	if err != nil {
		return nil, xerrors.Errorf("uc.repoUser.Get: %v", err)
	}
	return u, nil
}
