package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"library/database"
	"os"
	"time"
)

func isStudentInDB(userCred loginCred) bool {

	database.Connect()
	db := database.GetDBInstance()
	std := bson.D{
		{"Name", userCred.Name},
		{"Password", userCred.Password},
	}

	result := db.Collection("Students").FindOne(context.TODO(), std)

	if result.Err() == mongo.ErrNoDocuments {
		return false
	}
	return true
}

func generateTokens(userCred loginCred) tokens {

	var t tokens
	var err [2]error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": userCred.Name,
		"exp":  time.Now().Add(time.Minute * 15),
	})

	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": userCred.Name,
		"exp":  time.Now().Add(time.Hour * 72),
	})

	secretKeyAccessToken := []byte(os.Getenv("ACCESS_SECRET_KEY"))

	secretKeyRefreshToken := []byte(os.Getenv("REFRESH_SECRET_KEY"))

	// Sign and get the complete encoded token as a string using the secret
	t.AccessToken, err[0] = token.SignedString(secretKeyAccessToken)
	t.RefreshToken, err[1] = refToken.SignedString(secretKeyRefreshToken)

	return t
}

func VerifyAccessToken(accessToken string) bool {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		secretKeyAccessToken := []byte(os.Getenv("ACCESS_SECRET_KEY"))
		return secretKeyAccessToken, nil
	})
	if err != nil {
		fmt.Print(err)
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["name"], claims["exp"])
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
