package libroautorcontroller

import (
	"encoding/json"
	//"log"
	"net/http"

	//"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/models"
	DB "github.com/mejiadev643/config/db"
)

// var (
// 	validate *validator.Validate
// )

// func init() {
// 	validate = validator.New()
// }

func LibroAutorIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var libro []models.LibrosAutor
	DB.DB.Preload("Libro").Preload("Libro.Editorial").
		Preload("Libro.Genero").
		Preload("Autor").
		Where("deleted = ?", false).
		Find(&libro)
	json.NewEncoder(w).Encode(&libro)
}

func LibroAutorShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var libro []models.LibrosAutor
	params := mux.Vars(r) // se usa para traer todos los parametros
	//DB.DB.Where("deleted",false).First(&libro, params["id"])
	err := DB.DB.Preload("Libro").Preload("Libro.Editorial").
		Preload("Libro.Genero").
		Preload("Autores").
		Where("deleted = ?", false).
		Where("isbn = ?", params["isbn"]).
		Find(&libro).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Libro no encontrado")
		return
	}
	//hacer una consulta de autores segun los datos de libro y asignar a un nuevo objeto
	
    var sliceautor models.LibrosAutores
    sliceautor.Id = libro[0].Id
    sliceautor.ISBN = libro[0].ISBN
    sliceautor.LibroId = libro[0].LibroId
    sliceautor.Libro = libro[0].Libro
    sliceautor.Autores = []models.Autor{}
    for _, l := range libro {
        var autores models.Autor
        DB.DB.First(&autores, "id = ?", l.Autor.Id)
        sliceautor.Autores = append(sliceautor.Autores, autores)
    }
    json.NewEncoder(w).Encode(sliceautor)
}
