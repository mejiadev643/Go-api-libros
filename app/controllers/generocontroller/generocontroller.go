package generocontroller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/models"
	DB "github.com/mejiadev643/config/db"
	"net/http"
)

var (
    validate *validator.Validate
)

func init() {
    validate = validator.New()
}

func GeneroIndex(w http.ResponseWriter, r *http.Request) {
	var generos []models.Genero
	DB.DB.Where("deleted",false).Find(&generos)
	json.NewEncoder(w).Encode(&generos)

}

func GeneroShow(w http.ResponseWriter, r *http.Request) {
	var genero models.Genero
	params := mux.Vars(r) // se usa para traer todos los parametros
	DB.DB.Where("deleted",false).First(&genero, params["id"])
	if genero.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Genero no encontrado")
		return
	}
	json.NewEncoder(w).Encode(&genero)
}

func GeneroCreate(w http.ResponseWriter, r *http.Request) {
	var genero models.Genero
	json.NewDecoder(r.Body).Decode(&genero)

	err := validate.Struct(genero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	try := DB.DB.FirstOrCreate(&genero, genero)
	if try.Error != nil {
		json.NewEncoder(w).Encode(try.Error.Error())
		return
	}
	json.NewEncoder(w).Encode(&genero)
}

func GeneroUpdate(w http.ResponseWriter, r *http.Request) {
	var genero models.Genero
	param := mux.Vars(r)
	DB.DB.Where("deleted",false).First(&genero, param["id"])
	if genero.Id == 0 {
		json.NewEncoder(w).Encode("Genero no encontrado")
		return
	}

	var newGenero models.Genero
	err := json.NewDecoder(r.Body).Decode(&newGenero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newGenero.Id = genero.Id
	DB.DB.Model(&genero).Updates(newGenero)

	json.NewEncoder(w).Encode(&genero)
}

func GeneroDelete(w http.ResponseWriter, r *http.Request) {
    var genero models.Genero
    param := mux.Vars(r)
    DB.DB.First(&genero, param["id"])
    if genero.Id == 0 {
        json.NewEncoder(w).Encode("genero no encontrado")
        return
    }
    genero.Deleted = true
    DB.DB.Save(&genero)
    json.NewEncoder(w).Encode("genero eliminado")
}
