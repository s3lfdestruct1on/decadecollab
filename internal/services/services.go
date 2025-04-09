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