package handlers

import (
	"decadecollab/internal/models"
	"testing"
	"github.com/stretchr/testify/require"

	"github.com/go-playground/validator/v10"
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
				Title: "aa",
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
			require.Error(t, err)
		})
		
	}
	

}
