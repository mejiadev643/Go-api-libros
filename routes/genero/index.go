package genero

import (
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/controllers/generocontroller"
)

func Routes( r *mux.Router) {
	r.HandleFunc("", generocontroller.GeneroIndex).Methods("GET")
	r.HandleFunc("/", generocontroller.GeneroIndex).Methods("GET")
	r.HandleFunc("/", generocontroller.GeneroCreate).Methods("POST")
	r.HandleFunc("/{id}", generocontroller.GeneroShow).Methods("GET")
	r.HandleFunc("/{id}", generocontroller.GeneroUpdate).Methods("PUT")
	r.HandleFunc("/{id}", generocontroller.GeneroUpdate).Methods("DELETE")
}
