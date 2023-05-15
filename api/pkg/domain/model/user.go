package model

import (
	"github.com/google/uuid"
)

type User struct {
	id   string
	name string
}

func NewUser(id, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}

func NewUserID() string {
	return uuid.NewString()
}

func IsValidUserID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}
