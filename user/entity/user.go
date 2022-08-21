package entity

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" pg:"type:uuid"`
	Email    string    `json:"email"`
	password string
}

func CreateUser(e string, p string) *User {
	return &User{
		Id:       uuid.New(),
		Email:    e,
		password: p,
	}
}

type CreateUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
