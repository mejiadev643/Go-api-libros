package libroautor

import (
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/controllers/editorialcontroller"
	"github.com/mejiadev643/app/controllers/libroautorcontroller"
)

func Routes( r *mux.Router) {
	r.HandleFunc("", libroautorcontroller.LibroAutorIndex).Methods("GET")
	r.HandleFunc("/", libroautorcontroller.LibroAutorIndex).Methods("GET")
	r.HandleFunc("/", libroautorcontroller.LibroAutorShow).Methods("POST")
	r.HandleFunc("/{isbn}", libroautorcontroller.LibroAutorShow).Methods("GET")
	r.HandleFunc("/{isbn}", editorialcontroller.EditorialUpdate).Methods("PUT")
	r.HandleFunc("/{isbn}", editorialcontroller.EditorialDelete).Methods("DELETE")
}