package test_postgre

import (
	"context"
	"decadecollab/internal/config"
	"decadecollab/internal/lib/logger/sl"
	"decadecollab/internal/models"
	db "decadecollab/internal/storage"

	"github.com/go-playground/validator/v10"
)

var(
	sqlcfg = config.PSQLConfig{
	Port: "5432",
	Username: "admin",
	Password: "admin",
	Database: "my_db",
	}

    conn,_ = db.DBConn(sqlcfg)
    validate = validator.New()

)

func InitItemsDB() error {
    query := `
  		CREATE TABLE IF NOT EXISTS itemsTEMP (
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
        sl.PLogger.Error("unable to create itemsTEMP db", "error", err.Error())
    return err
    }
	return nil
}

func CreateItem(i *models.Item)(int64,error){
	err := validate.Struct(i)
	if err != nil {
		sl.PLogger.Error("invalid item info", "error", err.Error())
		return i.Id, err
	}
	
	query := "INSERT INTO itemsTEMP (title,stock,price,sale_percent,tags,description) values($1,$2,$3,$4,$5,$6) RETURNING id"

	err = conn.QueryRow(context.Background(),query,i.Title,i.Stock,i.Price,i.SalePercent,i.Tags,i.Description).Scan(&i.Id)
	if err != nil {
        sl.PLogger.Error("cannot create item", "error", err.Error())
		return i.Id, err
    }
	return i.Id,nil
}

func GetItem(id int64) (*models.Item,error){
	item := models.Item{}
	query := "SELECT * FROM itemsTEMP WHERE id = $1"

	err := conn.QueryRow(context.Background(),query,id).Scan(&item.Id,&item.Title,&item.Stock,&item.Price,&item.SalePercent,&item.Tags,&item.Description)
	if err != nil{
		sl.PLogger.Error("item not found", "error", err.Error())
		return &item, err
	}
	return &item,nil
}

func DropItem() (error){
	
	query := "DROP table itemsTEMP"

	_, err := conn.Exec(context.Background(),query)
	if err != nil{
		sl.PLogger.Error("unable to drop table", "error", err.Error())
		return err
	}
	return nil
}