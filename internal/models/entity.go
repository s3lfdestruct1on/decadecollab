package models

import (
  "time"
  "github.com/shopspring/decimal"
)
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
  Stock int16 
  Price decimal.Decimal
  SalePercent int16
  Tags string
  Description string
}

type Basket struct{
  Id int64
  UserID int64 
  ItemID int64
  Title string
  Quantity int16
  TotalPrice decimal.Decimal
}

type Order struct{
  Id int64
  User_id int64
  Total_price decimal.Decimal
  Shipping_address string
  Delivery_date time.Time
  
}

type Order_Items struct{
   Id int64
   Order_id int64
   Item_id int64
   Quantity int16
   Total_price decimal.Decimal
}