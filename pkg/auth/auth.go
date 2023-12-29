package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"library/pkg/common"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(userCred common.LoginCred) Tokens {

	var t Tokens
	var err [2]error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": userCred.Name,
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	})

	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": userCred.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})

	secretKeyAccessToken := []byte(os.Getenv("ACCESS_SECRET_KEY"))

	secretKeyRefreshToken := []byte(os.Getenv("REFRESH_SECRET_KEY"))

	// Sign and get the complete encoded token as a string using the secret
	t.AccessToken, err[0] = token.SignedString(secretKeyAccessToken)
	t.RefreshToken, err[1] = refToken.SignedString(secretKeyRefreshToken)
	if err[0] != nil && err[1] != nil {
		fmt.Errorf("tokens didn't get signed %v", err)
		return Tokens{}
	}
	return t
}

func VerifyAccessToken(accessToken string) (bool, string) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
	})
	if err != nil {

		fmt.Print("Error: ", err)
		return false, "" //common.LoginCred{}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//fmt.Println(claims["name"], claims["exp"])
		var name string //common.LoginCred
		name = claims["name"].(string)
		return true, name
	} else {
		fmt.Println("error:", err)
		return false, "" //common.LoginCred{}
	}
}

func VerifyRefreshToken(accessToken string) (bool, common.LoginCred) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("REFRESH_SECRET_KEY")), nil
	})
	if err != nil {

		fmt.Print("Error: ", err)
		return false, common.LoginCred{}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//fmt.Println(claims["name"], claims["exp"])
		var cred common.LoginCred
		cred.Name = claims["name"].(string)
		return true, cred
	} else {
		fmt.Println("error:", err)
		return false, common.LoginCred{}
	}
}

func GenerateAccessToken(userCred common.LoginCred) AccessToken {

	var t AccessToken
	var err error

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": userCred.Name,
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	})

	secretKeyAccessToken := []byte(os.Getenv("ACCESS_SECRET_KEY"))

	// Sign and get the complete encoded token as a string using the secret
	t.AccessToken, err = token.SignedString(secretKeyAccessToken)
	if err != nil {
		fmt.Errorf("tokens didn't get signed %v", err)
		return AccessToken{}
	}

	return t
}

func GetTokenFromRequest(c *gin.Context) (string, error) {
	var token string

	tokenHeader := c.GetHeader("Authorization")

	token = strings.TrimPrefix(tokenHeader, "Bearer ")
	if token != "" {
		return token, nil
	}
	return "", errors.New("something wrong with token string")
}
