package models

import "time"

// TODO: review всего этого дерьма сделать
type User struct {
	Id         int64
	Username   string `validate:"required,min=6,max=20,alphanum"`
	Password   string `validate:"required,min=4,max=25,alphanum"`
	Email      string `validate:"required,email"`
	LastActive time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Item struct {
	Id          int64  `json:"id"`
	Title       string `json:"title" validate:"min=3"`
	Stock       int64  `json:"stock" validate:"numeric,gt=0"`
	Price       int    `json:"price" validate:"numeric,gt=0"`
	SalePercent int16  `json:"saleperc"`
	Tags        string `json:"tags"`
	Description string `json:"desc"`
}

type Basket struct {
	Id         int64
	UserID     int64
	ItemID     int64
	TotalPrice int64
	Quantity   int16
}
