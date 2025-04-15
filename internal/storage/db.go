package db

import (
	"context"
	"decadecollab/internal/config"
	"decadecollab/internal/lib/logger/sl"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func DBConn(cfg config.PSQLConfig) (*pgx.Conn,error){
	
	
	usr := cfg.Username
	pass := cfg.Password
	db := cfg.Database
	port := cfg.Port
	strConn := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s",usr,pass,port,db)
	
	
	conn,err := pgx.Connect(context.Background(),strConn)
	if err != nil {
       sl.PLogger.Error("failed to connect to database","error", err.Error())
	   return nil, err	
    }
	return conn,nil
}