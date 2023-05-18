package repository

import (
	"context"
	"database/sql"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"golang.org/x/xerrors"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (repo *User) Get(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT * FROM users WHERE "id" = $1 LIMIT 1;`

	var uID, uName string
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&uID, &uName)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, xerrors.Errorf("repo.db.QueryRowContext: %w", ErrNotFound)
		default:
			return nil, xerrors.Errorf("repo.db.QueryRowContext: %v", err)
		}
	}
	return model.NewUser(uID, uName), nil
}

func (repo *User) Create(ctx context.Context, u *model.User) error {
	query := `INSERT INTO users("id","name") VALUES($1,$2);`

	_, err := repo.db.ExecContext(ctx, query, u.ID, u.Name)
	if err != nil {
		return xerrors.Errorf("repo.db.ExecContext: %v", err)
	}
	return nil
}

func (repo *User) Update(ctx context.Context, u *model.User) error {
	query := `UPDATE users SET "name"=$1 WHERE "id"=$2;`

	_, err := repo.db.ExecContext(ctx, query, u.Name, u.ID)
	if err != nil {
		return xerrors.Errorf("repo.db.ExecContext: %v", err)
	}
	return nil
}

func (repo *User) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE "id"=$1;`

	_, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return xerrors.Errorf("repo.db.ExecContext: %v", err)
	}
	return nil
}
