package routes

import (
	"github.com/gorilla/mux"
	"github.com/mejiadev643/routes/autor"
	"github.com/mejiadev643/routes/editorial"
	"github.com/mejiadev643/routes/genero"
	"github.com/mejiadev643/routes/libro"
	"github.com/mejiadev643/routes/libroautor"
	"github.com/mejiadev643/routes/login"
)

func RegisterRoutes( r *mux.Router) {//aqui se registran todas las rutas que se usaran en la aplicacion
	//se pueden separar por modulos para tener un mejor orden
	//nota, ya viene con el prefijo /api
	
	//rutas de login
	login.Routes(r)

	//ruta de genero
	genero.Routes(r.PathPrefix("/genero").Subrouter())
	
	//ruta de editorial
	editorial.Routes(r.PathPrefix("/editorial").Subrouter())

	//ruta de autor
	autor.Routes(r.PathPrefix("/autor").Subrouter())

	//ruta libro
	libro.Routes(r.PathPrefix("/libro").Subrouter())

	//ruta libroautor
	libroautor.Routes(r.PathPrefix("/libro-autor").Subrouter())

}