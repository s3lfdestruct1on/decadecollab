package fake_postgre

import (
	"decadecollab/internal/lib/logger/sl"
	"decadecollab/internal/models"
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
)


var(
	validate = validator.New()
	
)


func CreateItem(i *models.Item, i2 *models.Item)(error){
	err := validate.Struct(i)
	if err != nil {
		sl.PLogger.Error("invalid item info", "error", err.Error())
		return err
	}
	jsonmrh, err := json.Marshal(i)
	if err != nil {
        sl.PLogger.Error("cannot marshal item", "error", err.Error())
		return err
    }

	err = json.Unmarshal(jsonmrh, &i2)
	if err != nil {
        sl.PLogger.Error("cannot unmarshal item", "error", err.Error())
		return err
    }
	return nil
}

func GetItem(i *models.Item, id int64) (error){
	if i.Id != id{
		err := errors.New("i cho nahuy")
		sl.PLogger.Error("item not found", "error", err.Error())
		return err
	}
	return nil
}