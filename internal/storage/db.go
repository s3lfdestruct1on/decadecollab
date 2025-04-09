package db

import (
	"context"
	"decadecollab/internal/config"
	"fmt"

	"github.com/jackc/pgx/v5"
)

//logger sdelai :-)
func DBConn(cfg config.PSQLConfig) (*pgx.Conn,error){
	usr := cfg.Username
	pass := cfg.Password
	db := cfg.Database
	port := cfg.Port
	strConn := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s",usr,pass,port,db)
	
	
	conn,err := pgx.Connect(context.Background(),strConn)
	if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
	return conn,nil
}