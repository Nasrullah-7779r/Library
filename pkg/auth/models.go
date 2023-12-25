package auth

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type loginCred struct {
	Name     string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var validate *validator.Validate

func validateLoginCred(cred *loginCred) {

	validate = validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(cred)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
	}

}
