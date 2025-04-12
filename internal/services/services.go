package service

import (
	"decadecollab/internal/lib/logger/prettylog"
	"decadecollab/internal/models"
	"log/slog"

	postgre "decadecollab/internal/storage/sql"
	//"fmt"
)
var(
		loggatrr = slog.HandlerOptions{
		AddSource: true,
		}
		logg = slog.New(prettylog.NewHandler(&loggatrr))
)
// возвращает одного юзера по айди
func GetUser(id int)*models.User{
	if id <= 0{
		logg.Error("invalid id")
	}

	user,_ := postgre.GetUser(id)
	return user
}

func InitDB(){
	postgre.InitDB()
}
//создает пользователя с валидацей, смотреть условия
func NewUser(u *models.User)int{
	id,_ := postgre.CreateUser(u)
	return id
}

//удаляет юзера по айди id = user_id
func DeleteUSer(id int){
	postgre.DeleteUser(id)
}
// возвращает всех пользователй в диапазоне from - to, e.g. from 100 to 150 вернет 50 пользователей начиная с 100
func GetUsersFromTo(from int,to int)[]models.User{
	res,_ := postgre.GetUsersFromTo(from,to)
	return res
}
// возвращает страницу(10) пользователей
func GetUsersPaging(page int)[]models.User{
	res,_ := postgre.GetUsersPaging(page)
	return res
}
// апдейт с валидацией полностью пользователя,смотреть условия
func UpdateUser(id int,u *models.User){
	postgre.UpdateUser(id,u)
}
// смена только юзернейма с валидацией id = user_id, смотреть условия
func UpdateUsername(id int,username string){
	postgre.UpdateUsername(id,username)
}
// смена пароля с валидацией id = user_id
func UpdatePassword(id int,password string){
	postgre.UpdateUsername(id,password)
}
// смена только имейла с валидацией, id = user_id
func UpdateEmail(id int,email string){
	postgre.UpdateUsername(id,email)
}