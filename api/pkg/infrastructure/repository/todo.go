package repository

import (
	"context"
	"database/sql"

	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"golang.org/x/xerrors"
)

type Todo struct {
	db *sql.DB
}

func NewTodo(db *sql.DB) *Todo {
	return &Todo{
		db: db,
	}
}

func (repo *Todo) GetByUserID(ctx context.Context, id, userID string) (*model.Todo, error) {
	query := `SELECT "id","name","done" FROM "todos" WHERE "id" = $1 AND "user_id" = $2;`

	var t model.Todo
	err := repo.db.QueryRowContext(ctx, query, id, userID).Scan(&t.ID, &t.Name, &t.Done)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, xerrors.Errorf("repo.db.QueryRowContext: %w", ErrNotFound)
		default:
			return nil, xerrors.Errorf("repo.db.QueryRowContext: %v", err)
		}
	}
	return &t, nil
}

func (repo *Todo) GetAllByUserID(ctx context.Context, userID string) ([]*model.Todo, error) {
	query := `SELECT "id","name","done" FROM "todos" WHERE "user_id" = $1;`

	rows, err := repo.db.QueryContext(ctx, query, userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, xerrors.Errorf("repo.db.QueryRowContext: %w", ErrNotFound)
		default:
			return nil, xerrors.Errorf("repo.db.QueryRowContext: %v", err)
		}
	}
	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Name, &t.Done); err != nil {
			return nil, xerrors.Errorf("rows.Scan: %v",err)
		}
		todos = append(todos, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, xerrors.Errorf("rows.Err: %v",err)
	}
	return todos, nil
}
