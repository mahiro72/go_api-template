package usecase

import (
	"context"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type getUserWithTodo struct {
	repoUser repository.User
	repoTodo repository.Todo
}

func NewGetUserWithTodo(ru repository.User, rt repository.Todo) *getUserWithTodo {
	return &getUserWithTodo{
		repoUser: ru,
		repoTodo: rt,
	}
}

func (uc *getUserWithTodo) Exec(ctx context.Context, userID string) (*model.UserWithTodos, error) {
	todos, err := uc.repoTodo.GetAllByUserID(ctx, userID)
	if err != nil {
		return nil, xerrors.Errorf("uc.repoTodo.GetAllByUserID: %v", err)
	}
	u, err := uc.repoUser.Get(ctx, userID)
	if err != nil {
		return nil, xerrors.Errorf("uc.repoUser.Get: %v", err)
	}
	return model.NewUserWithTodos(u, todos), nil
}
