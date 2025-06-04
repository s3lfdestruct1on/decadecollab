package handlers

import (
	"decadecollab/internal/lib/logger/sl"
	"decadecollab/internal/models"
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
	id,err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil{
		sl.PLogger.Error("failed to parse id", "error", err.Error())
	}
	// GetUser возвращает models.User{} (обработка всех возможных ошибок происходит на стороне репозитория)
	user := service.GetUser(id)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error{
	id, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil{
		sl.PLogger.Error("user not found", "error", err.Error())
	}
	service.DeleteUser(id)
	return c.JSON(http.StatusOK,"user successfully deleted")
}

func CreateUser(c echo.Context) error{
 	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		sl.PLogger.Error("unable to create user", "error", err.Error())

	}
	service.NewUser(user)

	return c.JSON(http.StatusOK,"user successfully created")
}