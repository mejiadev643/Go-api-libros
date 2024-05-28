package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mejiadev643/app/models"
	"github.com/mejiadev643/config/db"
	"github.com/mejiadev643/database/seeders"
	"github.com/mejiadev643/routes"
)

func main() {
	DB.Connect()
	r:= mux.NewRouter()
	//migrar la base de datos
	DB.DB.AutoMigrate(&models.User{}, &models.Description{}, &models.Genero{}, &models.Editorial{}, &models.Autor{}, &models.LibrosAutor{})
	//seed
	seeders.Seed(DB.DB)
	//r.HandleFunc("/", routes.Hello).Methods("GET")
	api := r.PathPrefix("/api").Subrouter() //ruta base de la api, todas las rutas de la api deben tener /api
	routes.RegisterRoutes(api)// registro de rutas api

	log.Println("Servidor corriendo en http://localhost:8080")
	//servidor
	http.ListenAndServe(":8080", r)
}