package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/infrastructure/repository"
	"github.com/mahiro72/go_api-template/pkg/usecase"
)

type todoResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func newTodoResponse(t *model.Todo) *todoResponse {
	return &todoResponse{
		ID:   t.ID,
		Name: t.Name,
		Done: t.Done,
	}
}

func newTodosResponse(ts []*model.Todo) []*todoResponse {
	var r []*todoResponse
	for _, t := range ts {
		r = append(r, newTodoResponse(t))
	}
	return r
}

type userWithTodoResponse struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Todos []*todoResponse `json:"todos"`
}

func newUserWithTodoResponse(uwt *model.UserWithTodos) *userWithTodoResponse {
	return &userWithTodoResponse{
		ID:    uwt.ID,
		Name:  uwt.Name,
		Todos: newTodosResponse(uwt.Todos),
	}
}

func GetAllTodoByUserID(db *sql.DB) http.HandlerFunc {
	repoTodo := repository.NewTodo(db)
	repoUser := repository.NewUser(db)
	ucGetUserWithTodo := usecase.NewGetUserWithTodo(repoUser, repoTodo)

	handler := func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")

		uwt, err := ucGetUserWithTodo.Exec(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := newUserWithTodoResponse(uwt)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
	return handler
}

func GetTodoByUserID(db *sql.DB) http.HandlerFunc {
	repoTodo := repository.NewTodo(db)
	ucGetTodo := usecase.NewGetTodo(repoTodo)

	handler := func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		todoID := chi.URLParam(r, "todoID")

		todo, err := ucGetTodo.Exec(r.Context(), todoID, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := newTodoResponse(todo)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
	return handler
}
