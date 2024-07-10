package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mejiadev643/app/estucturas"
	DB "github.com/mejiadev643/config/db"
	"github.com/mejiadev643/app/models"
)

func CreateToken(user estucturas.Login) (string, error) {
	//Obtain user data
	println(user.Email)
	println(user.Password)
	var userData models.UserStruct
	DB.DB.Preload("Descriptions").Where("email = ? AND password = ?", user.Email, user.Password).First(&userData)

	// Create the Claims// equal to the payload
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = userData
	claims["permisions"] = []string{"read", "write", "delete"}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	// Print the token
	println(tokenString)
	return tokenString, nil
}
