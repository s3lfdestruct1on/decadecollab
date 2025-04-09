package main

import (
	"decadecollab/internal/config"
	"decadecollab/internal/server/handlers"
	db "decadecollab/internal/storage"
	postgre "decadecollab/internal/storage/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)
var (
	cfg = config.Config{
		AppHostPort: os.Getenv("HOSTPORT"),
	}
	sqlcfg = config.PSQLConfig{
		Port: os.Getenv("PORT"),
		Username: os.Getenv("USERDB"),
		Password: os.Getenv("PASS"),
		Database: os.Getenv("DB"),
	}
	
)
func main() {


	
	server := echo.New()

	apiv1 := server.Group("/api/v1")	

	apiv1.GET("/time", handlers.CurrentTime)

	if err := server.Start(cfg.AppHostPort); err != nil {
		fmt.Println(err.Error())
	}
	//TODO fatal sdelai :>
	
}
