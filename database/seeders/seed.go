package seeders

import (
	"log"
	"github.com/mejiadev643/app/models"
	"gorm.io/gorm"
)

// Seeder function
func Seed(db *gorm.DB) {
	// Crear datos de ejemplo para Genero
	generos := []models.Genero{
		{Id: 1, Nombre: "Ficción"},
		{Id: 2, Nombre: "No Ficción"},
	}

	for _, genero := range generos {
		if err := db.FirstOrCreate(&genero, genero).Error; err != nil {
			log.Printf("Error seeding genero: %v\n", err)
		}
	}

	// Crear datos de ejemplo para Editorial
	editoriales := []models.Editorial{
		{Id: 1, Nombre: "Editorial A"},
		{Id: 2, Nombre: "Editorial B"},
	}

	for _, editorial := range editoriales {
		if err := db.FirstOrCreate(&editorial, editorial).Error; err != nil {
			log.Printf("Error seeding editorial: %v\n", err)
		}
	}

	// Crear datos de ejemplo para Autor
	autores := []models.Autor{
		{Id: 1, Nombre: "Autor 1"},
		{Id: 2, Nombre: "Autor 2"},
	}

	for _, autor := range autores {
		if err := db.FirstOrCreate(&autor, autor).Error; err != nil {
			log.Printf("Error seeding autor: %v\n", err)
		}
	}
	//crear datos de ejemplo para Libro
	libros := []models.Libro{
		{ISBN: "978-3-16-148410-0", AñoPublicacion: "2021", NumeroPaginas: 200, GeneroId: 1, EditorialId: 1},
		{ISBN: "978-3-16-148411-7", AñoPublicacion: "2021", NumeroPaginas: 200, GeneroId: 2, EditorialId: 2},
	}
	for _, libro := range libros {
		if err := db.FirstOrCreate(&libro, libro).Error; err != nil {
			log.Printf("Error seeding libro: %v\n", err)
		}
	}

	// Crear datos de ejemplo para Libro
	librosAutores := []models.LibrosAutor{
		{ISBN: "978-3-16-148410-0", AutorId: 1, LibroId: "978-3-16-148410-0"},
		{ISBN: "978-3-16-148411-7", AutorId: 2, LibroId: "978-3-16-148411-7"},
	}

		for _, libroAutor := range librosAutores {
			if err := db.FirstOrCreate(&libroAutor, libroAutor).Error; err != nil {
				log.Printf("Error seeding libros_autor: %v\n", err)
			}
		}
}//find de la funcion Seed
