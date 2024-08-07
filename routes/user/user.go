package user

import (
	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/controllers/users"
	"github.com/mejiadev643/app/middleware"
)

func Routes(r *mux.Router) {
	r.Use(middleware.ControlMiddleware)
	r.HandleFunc("", users.UsersIndex).Methods("GET")
	r.HandleFunc("/", users.UsersIndex).Methods("GET")
	r.HandleFunc("/{id}", users.UsersShow).Methods("GET")
	r.HandleFunc("/", users.UsersCreate).Methods("POST")
	r.HandleFunc("/{id}", users.UsersUpdate).Methods("PUT")
	r.HandleFunc("/{id}", users.UsersDelete).Methods("DELETE")
}
