package models

import "time"
//TODO: review всего этого дерьма сделать
type User struct{
  Id int64
  Username string `validate:"required,min=6,max=20,alphanum"`
  Password string `validate:"required,min=4,max=25,alphanum"`
  Email string `validate:"required,email"`
  LastActive time.Time
  CreatedAt time.Time
  UpdatedAt time.Time
}

type Item struct{
  Id int64
  Title string 
  Stock int64 
  Price int
  SalePercent int16
  Tags string
  Description string
}

type Basket struct{
  Id int64
  UserID int64 
  ItemID int64 
  TotalPrice int64
  Quantity int16
}