package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mahiro72/go_api-template/pkg/domain/model"
	"github.com/mahiro72/go_api-template/pkg/infrastructure/repository"
	"github.com/mahiro72/go_api-template/pkg/usecase"
)

type userResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func newUserResponse(u *model.User) *userResponse {
	return &userResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}

func GetUser(db *sql.DB) http.HandlerFunc {
	repoUser := repository.NewUser(db)
	ucUser := usecase.NewGetUser(repoUser)

	handler := func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")

		user, err := ucUser.Exec(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := newUserResponse(user)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
	return handler
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	repoUser := repository.NewUser(db)
	ucUser := usecase.NewCreateUser(repoUser)

	handler := func(w http.ResponseWriter, r *http.Request) {
		var reqBody createUserRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := ucUser.Exec(r.Context(), reqBody.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := newUserResponse(user)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
	return handler
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	repoUser := repository.NewUser(db)
	ucUser := usecase.NewUpdateUser(repoUser)

	handler := func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")

		var reqBody updateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := ucUser.Exec(r.Context(), userID, reqBody.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "update ok")
		w.WriteHeader(http.StatusNoContent)
	}
	return handler
}

func DeleteUser(db *sql.DB) http.HandlerFunc {
	repoUser := repository.NewUser(db)
	ucUser := usecase.NewDeleteUser(repoUser)

	handler := func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")

		err := ucUser.Exec(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "delete ok")
		w.WriteHeader(http.StatusNoContent)
	}
	return handler
}

type createUserRequest struct {
	Name string `json:"name"`
}

type updateUserRequest struct {
	Name string `json:"name"`
}
