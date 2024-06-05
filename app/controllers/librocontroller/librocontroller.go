package librocontroller

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

func LibroIndex(w http.ResponseWriter, r *http.Request) {
	var libro []models.Libro
	DB.DB.Preload("Editorial").Preload("Genero").Where("deleted", false).Find(&libro)
	json.NewEncoder(w).Encode(&libro)

}

func LibroShow(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro
	params := mux.Vars(r) // se usa para traer todos los parametros
	isbn := params["isbn"]
	err := DB.DB.Preload("Editorial").Preload("Genero").Where("deleted", false).First(&libro, "isbn = ?", isbn).Error

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("libro no encontrado")
		return
	}
	json.NewEncoder(w).Encode(&libro)
}

func LibroCreate(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro
	json.NewDecoder(r.Body).Decode(&libro)

	err := validate.StructPartial(&libro, "ISBN", "Titulo", "AñoPublicacion", "NumeroPaginas", "GeneroId", "EditorialId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	try := DB.DB.FirstOrCreate(&libro, libro)
	if try.Error != nil {
		json.NewEncoder(w).Encode(try.Error.Error())
		return
	}
	// Crear la respuesta utilizando la estructura auxiliar para no mostrar relaciones
	libroResponse := models.LibroResponse{
		ISBN:           libro.ISBN,
		Titulo:         libro.Titulo,
		AñoPublicacion: libro.AñoPublicacion,
		NumeroPaginas:  libro.NumeroPaginas,
		GeneroId:       libro.GeneroId,
		EditorialId:    libro.EditorialId,
	}
	json.NewEncoder(w).Encode(&libroResponse)
}

func LibroUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var libro models.Libro
    param := mux.Vars(r)
    err := DB.DB.Where("deleted",false).First(&libro, "ISBN = ?", param["isbn"]).Error
    if err != nil {
        json.NewEncoder(w).Encode("libro no encontrado")
        return
    }

    var newlibro models.Libro
    er := json.NewDecoder(r.Body).Decode(&newlibro)
    if er != nil {
        http.Error(w, er.Error(), http.StatusBadRequest)
        return
    }

    newlibro.ISBN = libro.ISBN
    DB.DB.Model(&libro).Updates(newlibro)
	respuesta := models.LibroResponse{
		ISBN:           libro.ISBN,
		Titulo:         libro.Titulo,
		AñoPublicacion: libro.AñoPublicacion,
		NumeroPaginas:  libro.NumeroPaginas,
		GeneroId:       libro.GeneroId,
		EditorialId:    libro.EditorialId,
	}

    json.NewEncoder(w).Encode(respuesta)
}

func LibroDelete(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro
	param := mux.Vars(r)
	err := DB.DB.First(&libro, "isbn = ?",param["isbn"]).Error
	if err != nil{
	    json.NewEncoder(w).Encode("libro no encontrado")
	    return
	}
	libro.Deleted = true
	DB.DB.Save(&libro)
	json.NewEncoder(w).Encode("libro eliminado")
}
