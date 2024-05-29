package libro

import (
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/controllers/librocontroller"
)

func Routes( r *mux.Router) {
	r.HandleFunc("", librocontroller.LibroIndex).Methods("GET")
	r.HandleFunc("/", librocontroller.LibroIndex).Methods("GET")
	r.HandleFunc("/", librocontroller.LibroCreate).Methods("POST")
	r.HandleFunc("/{isbn}", librocontroller.LibroShow).Methods("GET")
	r.HandleFunc("/{isbn}", librocontroller.LibroUpdate).Methods("PUT")
	r.HandleFunc("/{isbn}", librocontroller.LibroDelete).Methods("DELETE")
}