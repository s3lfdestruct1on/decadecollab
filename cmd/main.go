package main

import (
	"decadecollab/internal/config"
	"decadecollab/internal/lib/logger/sl"
	"decadecollab/internal/models"
	"os"

	//postgre "decadecollab/internal/storage/sql"

	"decadecollab/internal/server/handlers"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)


var (
  cfg = config.Config{
	AppHostPort: os.Getenv("HOSTPORT"),
  }
)
func main() {

	
	
	server := echo.New()

	apiv1 := server.Group("/api/v1")	
	apiv1.GET("/time", handlers.CurrentTime)
	apiv1.GET("/user/:id", handlers.GetUser)
	if err := server.Start(cfg.AppHostPort); err != nil {
		sl.PLogger.Error("unable to start server", "error", err.Error())
	}
	
	
}
