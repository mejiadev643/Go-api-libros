package login

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mejiadev643/app/estucturas"
	"github.com/mejiadev643/app/jwt"
	DB "github.com/mejiadev643/config/db"
)

var (
	validate *validator.Validate
)
func init(){
	validate = validator.New()
}

func Login(w http.ResponseWriter, r *http.Request){
	var login estucturas.Login
	json.NewDecoder(r.Body).Decode(&login)
	err := validate.Struct(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//codificar el password con sha256
	hash := sha256.Sum256([]byte(login.Password))
	login.Password = hex.EncodeToString(hash[:])


	//consultar en la base de datos
	//si el usuario y contraseña son correctos
	//generar un token y devolverlo
	//si no son correctos devolver un mensaje de error
	var comparate estucturas.Login
	DB.DB.Where("email = ? AND password = ?", login.Email, login.Password).First(&comparate)
	if comparate.Email == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("El usuario no existe o la contraseña es incorrecta")
		return
	}
	if comparate.Email == login.Email && comparate.Password == login.Password {
		//json.NewEncoder(w).Encode("Token")
		//json.NewEncoder(w).Encode(comparate)
		token, _ := jwt.CreateToken(login)
		json.NewEncoder(w).Encode([]string{"token: ",token})
	}

}