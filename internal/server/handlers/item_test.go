package handlers

import (
	//"context"
	"decadecollab/internal/config"
	//"decadecollab/internal/lib/logger/sl"
	"decadecollab/internal/models"
	db "decadecollab/internal/storage"
	fake_postgre "decadecollab/internal/storage/fake_sql"
	"decadecollab/internal/storage/fake_sql/test_postgre"

	//postgre "decadecollab/internal/storage/sql"

	//postgre "decadecollab/internal/storage/sql"
	//"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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


func TestNewItem(t *testing.T) {

	model := &models.Item{
		Title: "aaaa",
		Stock: 1,
		Price: 1,
	}
	v := validator.New()
	err := v.Struct(model)
	require.NoError(t, err)

}

func TestErrorNewItem(t *testing.T) {

	cases := []struct {
		name  string
		model *models.Item
 	}{
		{
			name: "bad_item_title",
			model: &models.Item{
				Title: "aaa",
				Stock: 1,
				Price: 1,
			},
		},
		{
			name: "bad_stock_number",
			model: &models.Item{
				Title: "aaa",
				Stock: 0,
				Price: 1,
			},
		},
		{
			name: "bad_price_number",
			model: &models.Item{
				Title: "aaa",
				Stock: 1,
				Price: 0,
			},
		},
	}
	v := validator.New()
	for _, tCase := range cases{
		t.Run(tCase.name, func(t *testing.T){
			err := v.Struct(tCase.model)
			assert.Error(t, err)
		})
		
	}
	

}

//турбозаглушка
func TestTURBOCreateAndGetItem(t *testing.T){
	TMI := models.Item{}
	TNI := models.Item{
		Id: 1,
		Title: "Innuendo",
		Stock: 1488,
		Price: 1984,
	}
	err := fake_postgre.GetItem(&TMI, 0)
	require.NoError(t, err)

	err = fake_postgre.CreateItem(&TNI, &TMI)
	require.NoError(t, err)

	err = fake_postgre.GetItem(&TMI, TNI.Id)
	require.NoError(t, err)
}











func TestCreateAndGetItem(t *testing.T){
	ItemModel := models.Item{
		Title: "title",
		Stock: 1,
		Price: 1,
	}
	
	
	err := test_postgre.InitItemsDB()
	require.NoError(t, err)

	_, err = test_postgre.GetItem(1)
	//require.NoError(t, err)
	require.Equal(t, pgx.ErrNoRows, err)

	id, err := test_postgre.CreateItem(&ItemModel)
	require.NoError(t, err)

	_, err = test_postgre.GetItem(id)
	require.NoError(t, err)

	err = test_postgre.DropItem()
	require.NoError(t, err)
}