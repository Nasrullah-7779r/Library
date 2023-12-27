package auth

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"library/pkg/common"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

var validate *validator.Validate

func validateLoginCred(cred *common.LoginCred) {

	validate = validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(cred)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
	}

}
