package autor

import (
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/controllers/autorescontroller"
)

func Routes( r *mux.Router) {
	r.HandleFunc("", autorescontroller.AutorIndex).Methods("GET")
	r.HandleFunc("/", autorescontroller.AutorIndex).Methods("GET")
	r.HandleFunc("/", autorescontroller.AutorCreate).Methods("POST")
	r.HandleFunc("/{id}", autorescontroller.AutorShow).Methods("GET")
	r.HandleFunc("/{id}", autorescontroller.AutorUpdate).Methods("PUT")
	r.HandleFunc("/{id}", autorescontroller.AutorDelete).Methods("DELETE")
}