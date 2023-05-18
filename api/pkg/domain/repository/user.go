package repository

import (
	"context"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
)

type User interface {
	Get(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, u *model.User) error
	Update(ctx context.Context, u *model.User) error
	Delete(ctx context.Context, id string) error
}
