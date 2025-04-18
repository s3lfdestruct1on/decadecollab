package handlers

import (
	"decadecollab/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-playground/validator/v10"
)

func TestNewUer(t *testing.T) {

	model := &models.User{
		Username: "agfsegseg",
		Password: "fawdawd2",
		Email:    "adwdaw@awdawd.dwa",
	}
	v := validator.New()
	err := v.Struct(model)
	require.NoError(t, err)
}

func TestErrorNewUser(t *testing.T) {

	cases := []struct {
		name  string
		model *models.User
 	}{
		{
			name: "bad_username",
			model: &models.User{
				Username: "awdaawd124.com",
				Password: "awdse12412",
				Email: "wadwgdf@awd.awda",
			},
		},
		{
			name: "small_username",
			model: &models.User{
				Username: "awda",
				Password: "awdse12412",
				Email: "wadwgdf@awd.awda",
			},
		},
		{
			name: "bad_password",
			model: &models.User{
				Username: "wadawdadw",
				Password: "aws/adw",
				Email: "dawdawd@faw.com",
			},
		},
		{
			name: "small_password",
			model: &models.User{
				Username: "wadawdadw",
				Password: "aws",
				Email: "dawdawd@faw.com",
			},
		},
		{
			name: "bad_email",
			model: &models.User{
				Username: "awdawd12",
				Password: "dawdaw4",
				Email: "q2qe@adwaw",
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
