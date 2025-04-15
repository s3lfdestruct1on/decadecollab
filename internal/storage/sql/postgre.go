package postgre

import (
	"context"
	"decadecollab/internal/config"
	"decadecollab/internal/lib/logger/sl"
	"decadecollab/internal/models"
	db "decadecollab/internal/storage"
	"errors"
	"time"

	"os"

	"github.com/go-playground/validator/v10"
)


var(
	sqlcfg = config.PSQLConfig{
	Port: os.Getenv("PORT"),
	Username: os.Getenv("USERDB"),
	Password: os.Getenv("PASS"),
	Database: os.Getenv("DB"),
	}

   

    conn,_ = db.DBConn(sqlcfg)
    validate = validator.New()

)

func InitDB() error {
	query := `
	DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_type
            WHERE typname = 'rolename'
            AND   typtype = 'e'
        ) THEN
            CREATE TYPE rolename AS ENUM ('customer', 'manager', 'admin');
        END IF;
    END $$;
	
	CREATE TABLE IF NOT EXISTS users (
        id 			BIGSERIAL PRIMARY KEY,
        name 		VARCHAR(255) NOT NULL,
        password	VARCHAR(255) NOT NULL,
        email 		VARCHAR(255) UNIQUE NOT NULL,
		last_active TIMESTAMP,
		created_at 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		role 		ROLENAME DEFAULT 'customer'
    );`
    _, err := conn.Exec(context.Background(), query)
    if err != nil {
        sl.PLogger.Error("unable to init db at 3 part", "error", err.Error())
		return errors.New("gg")
    }
    
    return nil
}

func InitItemsDB() error {
    query := `
  		CREATE TABLE IF NOT EXISTS items (
        id 				BIGSERIAL PRIMARY KEY,
        title 			varchar(255) NOT NULL,
        price 			DECIMAL(10,2),
    	sale_percent 	smallint,
        stock 			SMALLINT,
    	description 	TEXT,
      	tags 			TEXT
    );`
    _, err := conn.Exec(context.Background(), query)
    if err != nil {
        sl.PLogger.Error("unable to create items db", "error", err.Error())
    return errors.New("gg")
    }
	return nil
}

func InitBasketDB() error {
    query := `
  		CREATE TABLE IF NOT EXISTS baskets (
        id 			BIGSERIAL PRIMARY KEY,
        user_id 	BIGINT NOT NULL,
        item_id 	BIGINT,
        quantity 	SMALLINT,
    	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
      	FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE
    );`
    _, err := conn.Exec(context.Background(), query)
    if err != nil {
        sl.PLogger.Error("failed to create basket db", "error", err.Error())
    return errors.New("gg")
    }
    
    return nil
}

func InitOrderItemsDB() error {
    query := `
  		CREATE TABLE IF NOT EXISTS order_items (
        id 			BIGINT PRIMARY KEY,
    	order_id 	BIGINT NOT NULL,
    	item_id 	BIGINT NOT NULL,
    	quantity 	INTEGER NOT NULL,
    	unit_price 	DECIMAL(10, 2) NOT NULL,
    	FOREIGN KEY (order_id) REFERENCES orders(id),
    	FOREIGN KEY (item_id) REFERENCES items(id)
    );`
    _, err := conn.Exec(context.Background(), query)
    if err != nil {
        sl.PLogger.Error("failed to create order_items db", "error", err.Error())
    return errors.New("gg")
    }
    return nil
}

func InitOrderDB() error {
	query := `
    DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_type
            WHERE typname = 'status'
            AND   typtype = 'e'
        ) THEN
            CREATE TYPE status AS ENUM ('created', 'payed', 'shipping','delivered');
        END IF;
    END $$;

    CREATE TABLE IF NOT EXISTS orders (
        id BIGSERIAL PRIMARY KEY,
        user_id BIGINT NOT NULL,
        total_amount DECIMAL(10, 2) NOT NULL,
        shipping_address TEXT NOT NULL,
        delivery_date TIMESTAMP,
        status status DEFAULT 'created',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );
    `
	_, err := conn.Exec(context.Background(), query)
    if err != nil {
        sl.PLogger.Error("failed to create order db", "error", err.Error())
		return errors.New("gg")
    }

    return nil
}


//создает пользователя с валидацей, смотреть условия
func CreateUser(u *models.User)(int,error){
	err := validate.Struct(u)
	if err != nil{
		sl.PLogger.Error("user is not valid", "error", err.Error())
		return u.Id, err
	}
	query := "INSERT INTO users (name,password,email) values($1,$2,$3) RETURNING id"
	
	err = conn.QueryRow(context.Background(),query,u.Username,u.Password,u.Email).Scan(&u.Id)
	if err != nil {
        sl.PLogger.Error("cannot create user", "error", err.Error())
		return u.Id, err
    }
	return u.Id,nil
}
// апдейт с валидацией полностью пользователя,смотреть условия
func UpdateUser(id int, u *models.User)error{
	var now time.Time

	err := validate.Struct(u)
	if err != nil{
		sl.PLogger.Error("user is not valid", "error", err.Error())
		return err
	}

	query := "UPDATE users SET name = $1,password = $2, email = $3,updated_at = $4 where id = $5"

	_,err = conn.Exec(context.Background(),query,u.Username,u.Password,u.Email,now,id)
	if err != nil {
        sl.PLogger.Error("cannot update user", "error", err.Error())		
		return err
    }
	return nil
}
// смена только юзернейма с валидацией id = user_id, смотреть условия
func UpdateUsername(id int,username string)error{
	var now time.Time

	err := validate.Var(username,"required,min=6,max=20,alphanum") //от 6 до 20 чаров, символы a-z A-Z 0-9
	if err != nil{
		sl.PLogger.Error("invalid username", "error", err.Error())
		return err
	}

	query := "UPDATE users SET name = $1,updated_at = $2 where id = $3"

	_,err = conn.Exec(context.Background(),query,username,now,id)
	if err != nil {
        sl.PLogger.Error("cannot update user", "error", err.Error())
		return err
    }
	return nil
}
// смена пароля с валидацией id = user_id
func UpdatePassword(id int,password string)error{
	var now time.Time

	err := validate.Var(password,"required,min=4,max=25,alphanum")
	if err != nil{
		sl.PLogger.Error("invalid password: must be 4-25 alphanumeric characters", "error", err.Error()) //от 4 до 25 чаров, сиволы a-z A-Z 0-9
	}

	query := "UPDATE users SET password = $1,updated_at = $2 where id = $3"

	_,err = conn.Exec(context.Background(),query,password,now,id)
	if err != nil {
        sl.PLogger.Error("cannot update password", "error", err.Error())
    }
	return nil
}

// смена только имейла с валидацией, id = user_id
func UpdateEmail(id int,email string)error{
	var now time.Time

	err := validate.Var(email,"required,email")
	if err != nil{
		sl.PLogger.Error("invalid email", "error", err.Error())
	}

	query := "UPDATE users SET email = $1, updated_at = $2 where id = $3"

	_,err = conn.Exec(context.Background(),query,email,now,id)
	if err != nil {
        sl.PLogger.Error("cannot update email", "error", err.Error())
    }
	return nil
}

// догадайся блять че делает 
func DeleteUser(id int)error{
	defer conn.Close(context.Background())
	query := "DELETE FROM users WHERE id = $1"

	_,err := conn.Exec(context.Background(),query,id)
	if err != nil{
		sl.PLogger.Error("user not found", "error", err.Error())
	}
	return nil
}
// возвращает одного юзера по айди
func GetUser(id int) (*models.User,error){
	user := models.User{}
	query := "SELECT * FROM users WHERE id = $1"

	err := conn.QueryRow(context.Background(),query,id).Scan(&user.Id,&user.Username,&user.Password,&user.Email)
	if err != nil{
		sl.PLogger.Error("user not found", "error", err.Error())
	}
	return &user,nil
}
// возвращает страницу(10) пользователей
func GetUsersPaging(page int)([]models.User,error){
	pagesize := 10
	offset := page-1
	var res []models.User
	query := "SELECT * FROM users limit $1 offset $2"

	rows,err := conn.Query(context.Background(),query,pagesize,offset)
	if err != nil{
		sl.PLogger.Error("failed to run query", "error", err.Error())
		return nil, err
	}

	for rows.Next(){
		user := models.User{}
		err := rows.Scan(&user.Id,&user.Username,&user.Password,&user.Email)
		if err != nil{
			sl.PLogger.Error("failed to scan user", "error", err.Error())
			return nil, err
		}
		res = append(res, user)
	}
	return res,nil
}
// возвращает всех пользователй в диапазоне from - to, e.g. from 100 to 150 вернет 50 пользователей начиная с 100
func GetUsersFromTo(from int,to int)([]models.User,error){
	limit := to - from
	var res []models.User
	query := "SELECT * FROM users limit $1 offset $2"

	rows,err := conn.Query(context.Background(),query,limit,from)
	if err != nil{
		sl.PLogger.Error("failed to run query", "error", err.Error())
		return nil, err
	}
	for rows.Next(){
		user := models.User{}
		err := rows.Scan(&user.Id,&user.Username,&user.Password,&user.Email)
		if err != nil{
			sl.PLogger.Error("failed to scan user", "error", err.Error())
			return nil, err
		}
		res = append(res, user)
	}
	return res,nil
}