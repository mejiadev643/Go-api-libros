package generocontroller

import (
	"encoding/json"
	"net/http"
	"github.com/mejiadev643/app/models"
	DB "github.com/mejiadev643/config/db"
	"github.com/gorilla/mux"
)

func GeneroIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var generos []models.Genero
	DB.DB.Find(&generos)
	json.NewEncoder(w).Encode(&generos)
	
}

func GeneroShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var genero models.Genero
	params := mux.Vars(r)// se usa para traer todos los parametros
	DB.DB.First(&genero, params["id"])
	if genero.Id == 0{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Genero no encontrado")
		return
	}
	json.NewEncoder(w).Encode(&genero)
}

func GeneroCreate(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    var genero models.Genero
    err := json.NewDecoder(r.Body).Decode(&genero)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validar los datos
    if genero.Nombre == "" {
        http.Error(w, "El nombre del género es requerido", http.StatusBadRequest)
        return
    }

    try := DB.DB.FirstOrCreate(&genero,genero)
    if try.Error != nil {
        json.NewEncoder(w).Encode(try.Error.Error())
        return
    }
    json.NewEncoder(w).Encode(&genero)
}

func GeneroUpdate(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    var genero models.Genero
    param := mux.Vars(r)
    DB.DB.First(&genero, param["id"])
    if genero.Id == 0{
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

func GeneroDelete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var genero models.Genero
	param := mux.Vars(r)
	DB.DB.First(&genero, param["id"])
	if genero.Id == 0{
		json.NewEncoder(w).Encode("Genero no encontrado")
		return
	}
	//actualmente se usa eliminacion fisica, se puede cambiar a eliminacion logica
	DB.DB.Delete(&genero)
	w.WriteHeader(http.StatusNoContent)
}