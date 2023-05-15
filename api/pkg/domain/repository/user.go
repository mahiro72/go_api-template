package repository

import (
	"context"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
)

type User interface {
	Get(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, id, name string) error
	Update(ctx context.Context, id, name string) error
	Delete(ctx context.Context, id string) error
}
