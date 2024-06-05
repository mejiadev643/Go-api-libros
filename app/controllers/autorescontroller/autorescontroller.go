package autorescontroller


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

func AutorIndex(w http.ResponseWriter, r *http.Request) {
	var autor []models.Autor
	DB.DB.Where("deleted", false).Find(&autor).Order("nombre ASC")
	json.NewEncoder(w).Encode(&autor)

}

func AutorShow(w http.ResponseWriter, r *http.Request) {
	var autor models.Autor
	params := mux.Vars(r) // se usa para traer todos los parametros
	DB.DB.Where("deleted",false).First(&autor, params["id"])
	if autor.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("autor no encontrado")
		return
	}
	json.NewEncoder(w).Encode(&autor)
}

func AutorCreate(w http.ResponseWriter, r *http.Request) {
	var autor models.Autor
	json.NewDecoder(r.Body).Decode(&autor)

	err := validate.Struct(autor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	try := DB.DB.FirstOrCreate(&autor, autor)
	if try.Error != nil {
		json.NewEncoder(w).Encode(try.Error.Error())
		return
	}
	json.NewEncoder(w).Encode(&autor)
}

func AutorUpdate(w http.ResponseWriter, r *http.Request) {
	var autor models.Autor
	param := mux.Vars(r)
	DB.DB.Where("deleted",false).First(&autor, param["id"])
	if autor.Id == 0 {
		json.NewEncoder(w).Encode("Autor no encontrado")
		return
	}

	var newAutor models.Autor
	err := json.NewDecoder(r.Body).Decode(&newAutor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newAutor.Id = autor.Id
	DB.DB.Model(&autor).Updates(newAutor)

	json.NewEncoder(w).Encode(&autor)
}

func AutorDelete(w http.ResponseWriter, r *http.Request) {
    var autor models.Autor
    param := mux.Vars(r)
    DB.DB.First(&autor, param["id"])
    if autor.Id == 0 {
        json.NewEncoder(w).Encode("autor no encontrado")
        return
    }
    autor.Deleted = true
    DB.DB.Save(&autor)
    json.NewEncoder(w).Encode("autor eliminado")
}