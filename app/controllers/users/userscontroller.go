package users

import (
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/models"
	DB "github.com/mejiadev643/config/db"
	"crypto/sha256"
    "encoding/hex"
)
var (
	validate *validator.Validate
)
func init(){
	validate = validator.New()
}

func UsersIndex(w http.ResponseWriter, r *http.Request){

	var users  []models.User

	DB.DB.Where("deleted", false).Find(&users)	
	json.NewEncoder(w).Encode(&users)
}

func UsersShow(w http.ResponseWriter, r *http.Request){
	var user models.UserStruct
	params := mux.Vars(r)
	DB.DB.Preload("Descriptions").Find(&user, params["id"])
	json.NewEncoder(w).Encode(&user)
}

func UsersCreate(w http.ResponseWriter, r *http.Request){
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    err := validate.StructPartial(&user, "Code","Email", "Nombres", "Password", "Super_password")
    if err != nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Codificar la contraseña con SHA256
    hash := sha256.Sum256([]byte(user.Password))
    user.Password = hex.EncodeToString(hash[:])

    // Codificar la super contraseña con SHA256
    hash = sha256.Sum256([]byte(user.Super_password))
    user.Super_password = hex.EncodeToString(hash[:])

    try := DB.DB.FirstOrCreate(&user, user)
    if try.Error != nil{
        json.NewEncoder(w).Encode(try.Error.Error())
        return
    }

    json.NewEncoder(w).Encode(&user)
}
func UsersUpdate(w http.ResponseWriter, r *http.Request){
	var user models.User
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&user)
	err := validate.StructPartial(&user, "Code","Email", "Nombres", "Password", "Super_password")
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Codificar la contraseña con SHA256
	hash := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])

	// Codificar la super contraseña con SHA256
	hash = sha256.Sum256([]byte(user.Super_password))
	user.Super_password = hex.EncodeToString(hash[:])

	DB.DB.Where("id = ?", params["id"]).Updates(&user)
	response := models.UserStruct{
		Id: user.Id,
		Code: user.Code,
		Email: user.Email,
		Nombres: user.Nombres,
	}
	json.NewEncoder(w).Encode(&response)
}

func UsersDelete(w http.ResponseWriter, r *http.Request){
	var user models.User
	params := mux.Vars(r)
	DB.DB.First(&user, params["id"])
	if user.Id == 0{
		json.NewEncoder(w).Encode("usuario no encontrado")
		return
	}
	user.Deleted = true
	DB.DB.Save(&user)
	json.NewEncoder(w).Encode("usuario eliminado")
}