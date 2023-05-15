package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (repo *User) Get(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT * FROM users WHERE "id" = $1 LIMIT 1;`

	var uID, uName string
	if err := repo.db.QueryRowContext(ctx, query, id).Scan(&uID, &uName); err != nil {
		switch err {

		}
		return nil, err
	}
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&uID, &uName)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, fmt.Errorf("User.Get: %w", ErrNotFound)
		default:
			return nil, err
		}
	}
	return model.NewUser(uID, uName), nil
}

func (repo *User) Create(ctx context.Context, id, name string) error {
	query := `INSERT INTO users("id","name") VALUES($1,$2);`

	_, err := repo.db.ExecContext(ctx, query, id, name)
	return err
}

func (repo *User) Update(ctx context.Context, id, name string) error {
	query := `UPDATE users SET "name"=$1 WHERE "id"=$2;`

	_, err := repo.db.ExecContext(ctx, query, id, name)
	return err
}

func (repo *User) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE "id"=$1;`

	_, err := repo.db.ExecContext(ctx, query, id)
	return err
}
