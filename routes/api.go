package routes

import (
	"github.com/gorilla/mux"
	"github.com/mejiadev643/routes/login"
	"github.com/mejiadev643/routes/genero"
)

func RegisterRoutes( r *mux.Router) {//aqui se registran todas las rutas que se usaran en la aplicacion
	//se pueden separar por modulos para tener un mejor orden
	//nota, ya viene con el prefijo /api
	
	//rutas de login
	login.Routes(r)

	//ruta de genero
	genero.Routes(r.PathPrefix("/genero").Subrouter())

}