package service

import (
	"context"
	"decadecollab/internal/config"
	"decadecollab/internal/models"
	db "decadecollab/internal/storage"
	postgre "decadecollab/internal/storage/sql"
	"fmt"
	"os"
)

var(
	sqlcfg = config.PSQLConfig{
	Port: os.Getenv("PORT"),
	Username: os.Getenv("USERDB"),
	Password: os.Getenv("PASS"),
	Database: os.Getenv("DB"),
}
)

func GetUserService(id int)*models.User{
	if id <= 0{
		fmt.Errorf("invalid id")
	}
	conn,err := db.DBConn(sqlcfg)
	if err != nil{
		fmt.Errorf("failed to connect to database: %w", err)
	}
	user,_ := postgre.GetUser(conn,id)
	defer conn.Close(context.Background())

	return user
}