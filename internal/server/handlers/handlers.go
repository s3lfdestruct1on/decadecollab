package handlers

import (
	postgre "decadecollab/internal/storage/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CurrentTime(c echo.Context) error {
  	return c.JSON(http.StatusOK, map[string]any{"current_time":time.Now().Format("2006-01-02 15:04:05")})
}

func GetUser(c echo.Context) error {
	// GetUser возвращает models.User{} (обработка всех возможных ошибок происходит на стороне репозитория)
	user := postgre.GetUser(c.Param("id"))
	return c.JSON(http.StatusOK, user)
}