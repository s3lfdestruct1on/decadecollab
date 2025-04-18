package handlers

import (
	"decadecollab/internal/lib/logger/sl"
	"decadecollab/internal/models"
	service "decadecollab/internal/services"
	//postgre "decadecollab/internal/storage/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CurrentTime(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{"current_time": time.Now().Format("2006-01-02 15:04:05")})
}

func GetUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		sl.PLogger.Error("user not found", "error", err.Error())
	}
	// GetUser возвращает models.User{} (обработка всех возможных ошибок происходит на стороне репозитория)
	user := service.GetUser(id)
	return c.JSON(http.StatusOK, user)
}

func NewItem(c echo.Context) error {
	NI := models.Item{}

	c.Bind(&NI)

	err := service.NewItem(&NI)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, NI)
}

func UpdateItem(c echo.Context) error {
	
	
	NI := models.Item{}

	c.Bind(&NI)

	_, err := service.GetItem(NI.Id)
	if err != nil{
		
		return c.JSON(http.StatusBadRequest, err)
	}

	err = service.UpdateItem(NI.Id, &NI)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, NI)
}
