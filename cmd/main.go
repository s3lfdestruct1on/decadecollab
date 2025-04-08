package main

import (
	"decadecollab/internal/config"
	"decadecollab/internal/server/handlers"
	"fmt"
	"os"
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

	if err := server.Start(cfg.AppHostPort); err != nil {
		fmt.Println(err.Error())
	}
}
