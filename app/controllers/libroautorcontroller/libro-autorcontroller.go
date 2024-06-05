package libroautorcontroller

import (
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/models"
	DB "github.com/mejiadev643/config/db"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func LibroAutorIndex(w http.ResponseWriter, r *http.Request) {
	var libro []models.LibrosAutor
	DB.DB.Preload("Libro").Preload("Libro.Editorial").
		Preload("Libro.Genero").
		Preload("Autor").
		Where("deleted = ?", false).
		Find(&libro)
	json.NewEncoder(w).Encode(&libro)
}

func LibroAutorShow(w http.ResponseWriter, r *http.Request) {
	var libros []models.LibrosAutor
	params := mux.Vars(r) // se usa para traer todos los parametros
	//DB.DB.Where("deleted",false).First(&libro, params["id"])
	err := DB.DB.Preload("Libro").Preload("Libro.Editorial").
		Preload("Libro.Genero").
		Preload("Autor").
		Where("deleted = ?", false).
		Where("isbn = ?", params["isbn"]).
		Find(&libros).Error
	if err != nil|| len(libros) == 0{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Libro no encontrado")
		return
	}
	
	if len(libros) == 1{
		json.NewEncoder(w).Encode(libros)
		return
	} else {
		var sliceautor models.LibrosAutores
		sliceautor.Id = libros[0].Id
		sliceautor.ISBN = libros[0].ISBN
		sliceautor.LibroId = libros[0].LibroId
		sliceautor.Libro = libros[0].Libro
		sliceautor.Autores = []models.Autor{}
		//hacer una consulta de autores segun los datos de libro y asignar a un nuevo objeto
		for _, l := range libros {
			var autores models.Autor
			DB.DB.First(&autores, "id = ?", l.Autor.Id)
			sliceautor.Autores = append(sliceautor.Autores, autores)
		}
		json.NewEncoder(w).Encode(sliceautor)
		return
	}

	
}

func LibroAutorCreate(w http.ResponseWriter, r *http.Request) {
	var libro models.LibrosAutor
	json.NewDecoder(r.Body).Decode(&libro)

	err := validate.StructPartial(&libro, "ISBN", "AutorId", "LibroId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	try := DB.DB.Preload("Libro").Preload("Libro.Editorial").Preload("Autor").FirstOrCreate(&libro, libro)
	if try.Error != nil {
		json.NewEncoder(w).Encode(try.Error.Error())
		return
	}
	json.NewEncoder(w).Encode(&libro)
}
func LibroAutorUpdate(w http.ResponseWriter, r *http.Request) {
	var libro models.LibrosAutor
	param := mux.Vars(r)
	DB.DB.Where("deleted",false).First(&libro, param["id"])
	if libro.Id == 0 {
		json.NewEncoder(w).Encode("Libro no encontrado")
		return
	}

	var newlibro models.LibrosAutor
	err := json.NewDecoder(r.Body).Decode(&newlibro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newlibro.Id = libro.Id
	DB.DB.Model(&libro).Updates(newlibro)

	json.NewEncoder(w).Encode(&libro)
}
func LibroAutorDelete(w http.ResponseWriter, r *http.Request) {
	var libro models.LibrosAutor
	param := mux.Vars(r)
	DB.DB.First(&libro, param["id"])
	if libro.Id == 0 {
		json.NewEncoder(w).Encode("Libro no encontrado")
		return
	}
	libro.Deleted = true
	DB.DB.Save(&libro)
	json.NewEncoder(w).Encode("relacion libro-autor eliminado")
}
