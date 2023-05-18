package model

import "github.com/google/uuid"

type Todo struct {
	ID   string
	Name string
	Done bool
}

func NewTodo(id, name string, done bool) *Todo {
	return &Todo{
		ID:   NewTodoID(),
		Name: name,
		Done: done,
	}
}

func NewTodoID() string {
	return uuid.NewString()
}

func IsValidTodoID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
