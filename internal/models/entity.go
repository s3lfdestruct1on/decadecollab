package models

import "time"

type User struct{
	Id int
	Username string `validate:"required,min=6,max=20,alphanum"`
	Password string `validate:"required,min=4,max=25,alphanum"`
	Email string `validate:"required,email"`
	LastActive time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}