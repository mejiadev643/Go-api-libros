package login

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/controllers/login"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("api login"))
}

func Routes( r *mux.Router) {
	r.HandleFunc("", Hello).Methods("GET")
	r.HandleFunc("/login", login.Login).Methods("POST")
}