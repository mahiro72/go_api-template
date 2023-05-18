package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

func NewUser(id, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}

func NewUserID() string {
	return uuid.NewString()
}

func IsValidUserID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

type UserWithTodos struct {
	ID    string
	Name  string
	Todos []*Todo
}

func NewUserWithTodos(u *User, todos []*Todo) *UserWithTodos {
	return &UserWithTodos{
		ID:    u.ID,
		Name:  u.Name,
		Todos: todos,
	}
}
