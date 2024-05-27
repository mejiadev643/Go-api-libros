package main

import (
	//"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mejiadev643/DB"
	"github.com/mejiadev643/models"
	"github.com/mejiadev643/routes"
	"github.com/mejiadev643/seeders"
)

func main() {
	DB.Connect()
	r:= mux.NewRouter()
	//migrar la base de datos
	DB.DB.AutoMigrate(&models.User{}, &models.Description{}, &models.Genero{}, &models.Editorial{}, &models.Autor{}, &models.LibrosAutor{})
	//seed
	seeders.Seed(DB.DB)
	r.HandleFunc("/", routes.Hello).Methods("GET")


	//servidor
	http.ListenAndServe(":8080", r)
}