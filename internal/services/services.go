package service

import (
	"decadecollab/internal/models"
	

	postgre "decadecollab/internal/storage/sql"
	"fmt"
)


func GetUser(id int)*models.User{
	if id <= 0{
		fmt.Errorf("invalid id")
	}

	user,_ := postgre.GetUser(id)
	return user
}

func InitDB(){
	postgre.InitDB()
}

func NewUser(username string,password string,email string)int{
	id,_ := postgre.CreateUser(username,password,email)
	return id
}


func DeleteUSer(id int){
	postgre.DeleteUser(id)
}

func GetUsersFromTo(from int,to int)[]models.User{
	res,_ := postgre.GetUsersFromTo(from,to)
	return res
}

func GetUsersPaging(page int)[]models.User{
	res,_ := postgre.GetUsersPaging(page)
	return res
}

func UpdateUser(id int,username string,password string,email string){
	postgre.UpdateUser(id,username,password,email)
}

func UpdateUsername(id int,username string){
	postgre.UpdateUsername(id,username)
}

func UpdatePassword(id int,password string){
	postgre.UpdateUsername(id,password)
}

func UpdateEmail(id int,email string){
	postgre.UpdateUsername(id,email)
}