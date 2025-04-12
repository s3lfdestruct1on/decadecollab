package main

import (
	"decadecollab/internal/config"
	"decadecollab/internal/lib/logger/prettylog"
	"decadecollab/internal/models"
	"log/slog"

	"decadecollab/internal/server/handlers"
	//"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)
var (
  cfg = config.Config{
    AppHostPort: os.Getenv("HOSTPORT"),
  }
  logerratr = slog.HandlerOptions{
	AddSource: true,
}
  logger = slog.New(prettylog.NewHandler(&logerratr))


  testuser = models.User{
	Id: 1,
	Username: "vxcvxxc",
	Password: "asdasdas",
	Email: "gg@gg.com",
  }
)
func main() {

	
	
	server := echo.New()

	apiv1 := server.Group("/api/v1")	
	apiv1.GET("/time", handlers.CurrentTime)
	apiv1.GET("/user/:id", handlers.GetUser)
	if err := server.Start(cfg.AppHostPort); err != nil {
		logger.Error("unable to start server", "error", err.Error())
	}
	
	
}
