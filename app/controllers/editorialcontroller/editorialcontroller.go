package editorialcontroller


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

func EditorialIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var editorial []models.Editorial
	DB.DB.Where("deleted", false).Find(&editorial).Order("nombre ASC")
	json.NewEncoder(w).Encode(&editorial)

}

func EditorialShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var editorial models.Editorial
	params := mux.Vars(r) // se usa para traer todos los parametros
	DB.DB.Where("deleted",false).First(&editorial, params["id"])
	if editorial.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Editorial no encontrado")
		return
	}
	json.NewEncoder(w).Encode(&editorial)
}

func EditorialCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var editorial models.Editorial
	json.NewDecoder(r.Body).Decode(&editorial)

	err := validate.Struct(editorial)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	try := DB.DB.FirstOrCreate(&editorial, editorial)
	if try.Error != nil {
		json.NewEncoder(w).Encode(try.Error.Error())
		return
	}
	json.NewEncoder(w).Encode(&editorial)
}

func EditorialUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var editorial models.Editorial
	param := mux.Vars(r)
	DB.DB.Where("deleted",false).First(&editorial, param["id"])
	if editorial.Id == 0 {
		json.NewEncoder(w).Encode("editorial no encontrado")
		return
	}

	var neweditorial models.Editorial
	err := json.NewDecoder(r.Body).Decode(&neweditorial)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	neweditorial.Id = editorial.Id
	DB.DB.Model(&editorial).Updates(neweditorial)

	json.NewEncoder(w).Encode(&editorial)
}

func EditorialDelete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var editorial models.Editorial
    param := mux.Vars(r)
    DB.DB.First(&editorial, param["id"])
    if editorial.Id == 0 {
        json.NewEncoder(w).Encode("editorial no encontrado")
        return
    }
    editorial.Deleted = true
    DB.DB.Save(&editorial)
    json.NewEncoder(w).Encode("editorial eliminado")
}