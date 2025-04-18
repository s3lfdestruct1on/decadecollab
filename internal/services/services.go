package service

import (
	"decadecollab/internal/lib/logger/sl"
	"decadecollab/internal/models"

	postgre "decadecollab/internal/storage/sql"

	"github.com/go-playground/validator/v10"
	//"fmt"
)

var (
	validate = validator.New()
)

// возвращает одного юзера по айди
func GetUser(id int64) *models.User {
	if id <= 0 {
		sl.PLogger.Error("invalid id")
	}

	user, _ := postgre.GetUser(id)
	return user
}

/*
//Вовращает итем по айди
func GetItem(id int64)*models.Item{
	if id <= 0{
		sl.PLogger.Error("invalid id")
	}

	item,_ := postgre.GetItem(id)
	return item
}
*/

func InitDB() {
	postgre.InitDB()
}

// создает пользователя с валидацей, смотреть условия
func NewUser(u *models.User) int64 {
	id, _ := postgre.CreateUser(u)
	return id
}

// удаляет юзера по айди id = user_id
func DeleteUser(id int64) {
	postgre.DeleteUser(id)
}

// возвращает всех пользователй в диапазоне from - to, e.g. from 100 to 150 вернет 50 пользователей начиная с 100
func GetUsersFromTo(from int, to int) []models.User {
	res, _ := postgre.GetUsersFromTo(from, to)
	return res
}

// возвращает страницу(10) пользователей
func GetUsersPaging(page int) []models.User {
	res, _ := postgre.GetUsersPaging(page)
	return res
}

// апдейт с валидацией полностью пользователя,смотреть условия
func UpdateUser(id int64, u *models.User) {
	postgre.UpdateUser(id, u)
}

// смена только юзернейма с валидацией id = user_id, смотреть условия
func UpdateUsername(id int64, username string) {
	postgre.UpdateUsername(id, username)
}

// смена пароля с валидацией id = user_id
func UpdatePassword(id int64, password string) {
	postgre.UpdateUsername(id, password)
}

// смена только имейла с валидацией, id = user_id
func UpdateEmail(id int64, email string) {
	postgre.UpdateUsername(id, email)
}

func InitItemsDB() {
	postgre.InitItemsDB()
}

// Создаёт итем из модели models.item, возвращая айди итема или ошибку
func NewItem(i *models.Item) (error){
	_, err := postgre.CreateItem(i)
	if err != nil{
		return err
	}
	return nil
}


func GetItem(id int64) (*models.Item, error) {
	GI := &models.Item{}
	GI, err := postgre.GetItem(id)
	if err != nil{
		sl.PLogger.Error("invalid item info", "error", err.Error())
		return GI, err
	}
	return GI, nil
}


// Полный апдейт итема при помощи модели models.item по айди, возвращает, или ноль при удаче, или ошибку
func UpdateItem(id int64, i *models.Item) error {
	
	
	err := postgre.UpdateItem(id, i)
	if err != nil{
		sl.PLogger.Error("invalid item info", "error", err.Error())
		return err
	}
	return nil
}

// Меняет название итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemTitle(id int64, itemtitle string) {
	postgre.UpdateItemTitle(id, itemtitle)
}

// Меняет количесвто на складе итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemStock(id int64, itemstock int64) {
	postgre.UpdateItemStock(id, itemstock)
}

// Меняет цену итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemPrice(id int64, itemprice int) {
	postgre.UpdateItemPrice(id, itemprice)
}

// Меняет процент сккидки итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemSalePerc(id int64, itemsaleperc int16) {
	postgre.UpdateItemSalePerc(id, itemsaleperc)
}

// Меняет теги итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemTags(id int64, itemtags string) {
	postgre.UpdateItemTags(id, itemtags)
}

// Меняет описание итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemDesc(id int64, itemdesc string) {
	postgre.UpdateItemDesc(id, itemdesc)
}
