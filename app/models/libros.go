package models

type Libro struct {
	ISBN           string        `gorm:"column:isbn;primaryKey" json:"isbn"`
	AñoPublicacion string        `gorm:"column:año_publicacion" json:"año_publicacion"`
	NumeroPaginas  int           `gorm:"column:numero_paginas" json:"numero_paginas"`
	GeneroId       uint          `gorm:"column:genero_id;foreignKey:GeneroId" json:"id_genero"`
	EditorialId    uint          `gorm:"column:editorial_id;foreignKey:EditorialId" json:"editorial"`
	LibrosAutor    []LibrosAutor `gorm:"foreignKey:LibroId"`
	Editorial      Editorial     `gorm:"foreignKey:EditorialId"`
}

func (Libro) TableName() string {
	return "libros"
}
