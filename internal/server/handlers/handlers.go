package handlers

import (
	"decadecollab/internal/lib/logger/prettylog"
	service "decadecollab/internal/services"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)
var(
		loggatrr = slog.HandlerOptions{
		AddSource: true,
		}
		logg = slog.New(prettylog.NewHandler(&loggatrr))
)
func CurrentTime(c echo.Context) error {
  	return c.JSON(http.StatusOK, map[string]any{"current_time":time.Now().Format("2006-01-02 15:04:05")})
}

func GetUser(c echo.Context) error {
	id,err :=strconv.Atoi(c.Param("id"))
	if err != nil{
		logg.Error("user not found", "error", err.Error())
	}
	// GetUser возвращает models.User{} (обработка всех возможных ошибок происходит на стороне репозитория)
	user := service.GetUser(id)
	return c.JSON(http.StatusOK, user)
}