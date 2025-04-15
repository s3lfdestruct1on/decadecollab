package handlers

import (
	"decadecollab/internal/lib/logger/sl"
	service "decadecollab/internal/services"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CurrentTime(c echo.Context) error {
  	return c.JSON(http.StatusOK, map[string]any{"current_time":time.Now().Format("2006-01-02 15:04:05")})
}

func GetUser(c echo.Context) error {
	id,err :=strconv.Atoi(c.Param("id"))
	if err != nil{
		sl.PLogger.Error("user not found", "error", err.Error())
	}
	// GetUser возвращает models.User{} (обработка всех возможных ошибок происходит на стороне репозитория)
	user := service.GetUser(id)
	return c.JSON(http.StatusOK, user)
}