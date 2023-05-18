package repository

import (
	"context"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
)

type Todo interface {
	GetByUserID(ctx context.Context, id, userID string) (*model.Todo, error)
	GetAllByUserID(ctx context.Context, userID string) ([]*model.Todo, error)
}
