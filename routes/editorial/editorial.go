package editorial

import (
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/controllers/editorialcontroller"
)

func Routes( r *mux.Router) {
	r.HandleFunc("", editorialcontroller.EditorialIndex).Methods("GET")
	r.HandleFunc("/", editorialcontroller.EditorialIndex).Methods("GET")
	r.HandleFunc("/", editorialcontroller.EditorialCreate).Methods("POST")
	r.HandleFunc("/{id}", editorialcontroller.EditorialShow).Methods("GET")
	r.HandleFunc("/{id}", editorialcontroller.EditorialUpdate).Methods("PUT")
	r.HandleFunc("/{id}", editorialcontroller.EditorialDelete).Methods("DELETE")
}