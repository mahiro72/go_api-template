package usecase

import (
	"context"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type getTodo struct {
	repoTodo repository.Todo
}

func NewGetTodo(rt repository.Todo) *getTodo {
	return &getTodo{repoTodo: rt}
}

func (uc *getTodo) Exec(ctx context.Context, id, userID string) (*model.Todo, error) {
	if !model.IsValidTodoID(id) {
		return nil, xerrors.Errorf("!model.IsValidTodoID: id is invalid")
	}
	if !model.IsValidUserID(userID) {
		return nil, xerrors.Errorf("!model.IsValidUserID: userID is invalid")
	}
	t, err := uc.repoTodo.GetByUserID(ctx, id, userID)
	if err != nil {
		return nil, xerrors.Errorf("uc.repoTodo.GetByUserID: %v", err)
	}
	return t, nil
}
