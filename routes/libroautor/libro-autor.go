package libroautor

import (
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/controllers/libroautorcontroller"
)

func Routes( r *mux.Router) {
	r.HandleFunc("", libroautorcontroller.LibroAutorIndex).Methods("GET")
	r.HandleFunc("/", libroautorcontroller.LibroAutorIndex).Methods("GET")
	r.HandleFunc("/", libroautorcontroller.LibroAutorCreate).Methods("POST")
	r.HandleFunc("/{isbn}", libroautorcontroller.LibroAutorShow).Methods("GET")
	r.HandleFunc("/{isbn}", libroautorcontroller.LibroAutorUpdate).Methods("PUT")
	r.HandleFunc("/{isbn}", libroautorcontroller.LibroAutorDelete).Methods("DELETE")
}