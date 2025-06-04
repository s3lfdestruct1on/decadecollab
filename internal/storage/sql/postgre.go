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
    "github.com/shopspring/decimal"
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
    	sale_percent 	SMALLINT,
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
        id 			BIGSERIAL PRIMARY KEY,
    	order_id 	BIGINT NOT NULL,
    	item_id 	BIGINT NOT NULL,
    	quantity 	INTEGER NOT NULL,
    	total_price DECIMAL(10, 2) NOT NULL,
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
        total_price DECIMAL(10, 2) NOT NULL,
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
func CreateUser(u *models.User)(int64,error){
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
func UpdateUser(id int64, u *models.User)error{
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
func UpdateUsername(id int64,username string)error{
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
func UpdatePassword(id int64,password string)error{
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
func UpdateEmail(id int64,email string)error{
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
func DeleteUser(id int64)error{
	query := "DELETE FROM users WHERE id = $1"

	_,err := conn.Exec(context.Background(),query,id)
	if err != nil{
		sl.PLogger.Error("user not found", "error", err.Error())
	}
	return nil
}
// возвращает одного юзера по айди
func GetUser(id int64) (*models.User,error){
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
	defer rows.Close()
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
	defer rows.Close()
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
//Создаёт итем из модели models.item, возвращая айди итема или ошибку
func CreateItem(i *models.Item)(int64,error){
	query := "INSERT INTO items (title,stock,price,sale_percent,tags,description) values($1,$2,$3,$4,$5,$6) RETURNING id"

	err := conn.QueryRow(context.Background(),query,i.Title,i.Stock,i.Price.String(),i.SalePercent,i.Tags,i.Description).Scan(&i.Id)
	if err != nil {
        sl.PLogger.Error("cannot create item", "error", err.Error())
		return i.Id, err
    }
	return i.Id,nil
}
//Полный апдейт итема при помощи модели models.item по айди, возвращает, или ноль при удаче, или ошибку
func UpdateItem(id int64, i *models.Item)error{
	query := "UPDATE items SET title = $1,stock = $2,price = $3,sale_percent = $4,tags = $5, description = $6 where id = $7"

	_, err := conn.Exec(context.Background(),query,i.Title,i.Stock,i.Price.String(),i.SalePercent,i.Tags,i.Description,id)
	if err != nil{
		sl.PLogger.Error("cannot update item", "error", err.Error())
		return err
	}
	return nil

}


//Меняет название итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemTitle(id int64, itemtitle string)error{

 query := "UPDATE items SET title = $1 where id = $2"

 _,err := conn.Exec(context.Background(),query,itemtitle,id)
	if err != nil {
        sl.PLogger.Error("cannot update item", "error", err.Error())
		return err
    }
	return nil

}
//Меняет количесвто на складе итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemStock(id int64, itemstock int16)error{

	query := "UPDATE items SET stock = $1 where id = $2"
   
	_,err := conn.Exec(context.Background(),query,itemstock,id)
	   if err != nil {
		   sl.PLogger.Error("cannot update item", "error", err.Error())
		   return err
	   }
	   return nil
   
   }
//Меняет цену итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemPrice(id int64, itemprice decimal.Decimal)error{

 query := "UPDATE items SET price = $1 where id = $2"

 _,err := conn.Exec(context.Background(),query,itemprice.String(),id)
	if err != nil {
        sl.PLogger.Error("cannot update item", "error", err.Error())
		return err
    }
	return nil

}   
//Меняет процент сккидки итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemSalePerc(id int64, itemsaleperc int16)error{

	query := "UPDATE items SET sale_percent = $1 where id = $2"
   
	_,err := conn.Exec(context.Background(),query,itemsaleperc,id)
	   if err != nil {
		   sl.PLogger.Error("cannot update item", "error", err.Error())
		   return err
	   }
	   return nil
   
}
//Меняет теги итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemTags(id int64, itemtags string)error{

	query := "UPDATE items SET tags = $1 where id = $2"
   
	_,err := conn.Exec(context.Background(),query,itemtags,id)
	   if err != nil {
		   sl.PLogger.Error("cannot update item", "error", err.Error())
		   return err
	   }
	   return nil
   
}
//Меняет описание итема по айдишнику, возвращает ошибку, либо ноль при удаче
func UpdateItemDesc(id int64, itemdesc string)error{

	query := "UPDATE items SET description = $1 where id = $2"
   
	_,err := conn.Exec(context.Background(),query,itemdesc,id)
	   if err != nil {
		   sl.PLogger.Error("cannot update item", "error", err.Error())
		   return err
	   }
	   return nil
}
//добавляет товар по item_id в корзину пользователя по user_id с проверкой стока 
func AddItemToBasket(item_id int64,user_id int64)error{
	tx, err := conn.Begin(context.Background())
	defer tx.Rollback(context.Background())
    
    var stock int16
    err = tx.QueryRow(context.Background(), `
        SELECT stock
        FROM items
        WHERE id = $1
        FOR UPDATE;
    `, item_id).Scan(&stock)
	if stock < 1{
		sl.PLogger.Error("insufficient stock")
		return err
	}

    if err != nil {
        sl.PLogger.Error("failed to check item availability", err.Error())
		return err
    }

    _, err = tx.Exec(context.Background(), `
        INSERT INTO baskets (user_id, item_id, quantity)
        VALUES ($1, $2, 1)
        ON CONFLICT (user_id, item_id) DO UPDATE
        SET quantity = baskets.quantity + 1;
    `, user_id, item_id)

    if err != nil {
        sl.PLogger.Error("failed to add item to basket", err.Error())
		return err
    }

    err = tx.Commit(context.Background())
    if err != nil {
        sl.PLogger.Error("failed to commit transaction",err.Error())
		return err
    }

    sl.PLogger.Error("Added item %d to basket for user %d\n", item_id, user_id)
    return nil
}

func RemoveItemFromBasket(user_id int64,item_id int64)error{
    tx, err := conn.Begin(context.Background())
    if err != nil {
        sl.PLogger.Error("failed to begin transaction: %w", err)
        return err
    }
    defer tx.Rollback(context.Background()) 

    var quantity int
    err = tx.QueryRow(context.Background(), `
        SELECT quantity
        FROM baskets
        WHERE user_id = $1 AND item_id = $2
        FOR UPDATE;
    `, user_id, item_id).Scan(&quantity)

    if err != nil {
        sl.PLogger.Error("failed to check item in basket: %w", err)
        return err
    }
    if quantity > 1 {
         _, err = tx.Exec(context.Background(), `
            UPDATE baskets
            SET quantity = quantity - 1
            WHERE user_id = $1 AND item_id = $2;
        `, user_id, item_id)
    } else {
        _, err = tx.Exec(context.Background(), `
            DELETE FROM baskets
            WHERE user_id = $1 AND item_id = $2;
        `, user_id, item_id)
        }
     

    if err != nil {
        sl.PLogger.Error("failed to update or delete item from basket: %w", err)
        return err
	}

    err = tx.Commit(context.Background())
    if err != nil {
        sl.PLogger.Error("failed to commit transaction: %w", err)
        return err
    }

    sl.PLogger.Info("Removed item %d from basket for user %d", item_id, user_id)
    return nil
}

func GetBasketItems(user_id int64)([]models.Basket,error){
	query := `
        SELECT
            b.id,
			b.user_id,
            b.item_id, 
            i.title,
            b.quantity,
            (i.price * b.quantity) AS total_price
        FROM baskets b
        JOIN items i ON b.item_id = i.id
        WHERE b.user_id = $1;
    `

    rows, err := conn.Query(context.Background(), query, user_id)
    if err != nil {
        sl.PLogger.Error("failed to fetch basket items: %w", err)
        return nil, err
    }
    defer rows.Close()

    var basketItems []models.Basket
    for rows.Next() {
        var basket models.Basket
        err := rows.Scan(
            &basket.Id,
			&basket.UserID,
            &basket.ItemID,
            &basket.Title,
            &basket.Quantity,
            &basket.TotalPrice,
        )
        if err != nil {
            sl.PLogger.Error("failed to scan basket item: %w", err)
            return nil,err
        }
        basketItems = append(basketItems, basket)
    }

    return basketItems, nil
}

func CreateOrdersFromBaskets(user_id int64,shipping_address string,delivery_date time.Time)error{
	tx, err := conn.Begin(context.Background())
    if err != nil {
        sl.PLogger.Error("failed to begin transaction: %w", err)
        return err
    }
    defer tx.Rollback(context.Background()) 
	baskets,err:= GetBasketItems(user_id)
	if err != nil{
		sl.PLogger.Error("failed to get baskets: %w", err)
	}

	var order_id int64
	err = tx.QueryRow(context.Background(), `
        INSERT INTO orders (user_id, shipping_address,delivery_date,total_price)
        VALUES ($1, $2, $3,$4)
        RETURNING id
    `, user_id, shipping_address,delivery_date, 1).Scan(&order_id)
    if err != nil {
		sl.PLogger.Error("failed to create order: %w", err)
        return err
    }
    var total_price decimal.Decimal
	for _,basket := range baskets{
		_, err = tx.Exec(context.Background(), `
        INSERT INTO order_items (order_id, item_id, quantity, total_price)
        VALUES($1, $2, $3, $4)
    `, order_id, basket.ItemID,basket.Quantity,basket,basket.TotalPrice)
    if err != nil {
        sl.PLogger.Error("failed to create order_items: %w", err)
        return err
    }
    total_price.Add(basket.TotalPrice)
	}
	_, err = tx.Exec(context.Background(), `
        UPDATE orders
        SET total_price = $1
        WHERE id = $2
    `, total_price.String(), order_id)
    if err != nil {
        sl.PLogger.Error("failed to update order total price: %w", err)
        return err
    }
    _, err = tx.Exec(context.Background(), `
        DELETE FROM baskets
        WHERE user_id = $1
    `, user_id)
    if err != nil {
        sl.PLogger.Error("failed to delete  order baskets: %w", err)
        return err
    }
    err = tx.Commit(context.Background()) 
    if err != nil {
        sl.PLogger.Error("failed to commit: %w", err)
        return err
    }
    return nil
}