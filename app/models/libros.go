package models

import (
	"encoding/json"
	"time"
)

type Libro struct {
	ISBN           string    `gorm:"column:isbn;primaryKey" json:"isbn" validate:"required"`
	Titulo         string    `gorm:"column:titulo" json:"titulo" validate:"required"`
	AñoPublicacion time.Time    `gorm:"column:año_publicacion;type:date" json:"año-publicacion" validate:"required"`
	NumeroPaginas  int       `gorm:"column:numero_paginas" json:"numero-paginas" validate:"required"`
	GeneroId       uint      `gorm:"column:genero_id;foreignKey:GeneroId" json:"genero-id" validate:"required"`
	EditorialId    uint      `gorm:"column:editorial_id;foreignKey:EditorialId" json:"editorial-id" validate:"required"`
	Editorial      Editorial `gorm:"foreignKey:EditorialId" json:"editorial"`
	Genero         Genero    `gorm:"foreignKey:GeneroId" json:"genero"`
	Deleted        bool      `gorm:"column:deleted;default:false" json:"-" `
}

type LibroResponse struct {
	ISBN           string `json:"isbn"`
	Titulo         string `json:"titulo"`
	AñoPublicacion time.Time `json:"año_publicacion"`
	NumeroPaginas  int    `json:"numero_paginas"`
	GeneroId       uint   `json:"genero_id"`
	EditorialId    uint   `json:"editorial_id"`
}

func (Libro) TableName() string {
	return "libros"
}
func (a *Libro) UnmarshalJSON(data []byte) error {
    type Alias Libro
    aux := &struct {
        AñoPublicacion string `json:"año-publicacion"`
        *Alias
    }{
        Alias: (*Alias)(a),
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    t, err := time.Parse("2006", aux.AñoPublicacion)
    if err != nil {
        return err
    }

    a.AñoPublicacion = t

    return nil
}

func (a *Libro) MarshalJSON() ([]byte, error) {
    type Alias Libro
    return json.Marshal(&struct {
        *Alias
        AñoPublicacion string `json:"año-publicacion"`
    }{
        Alias:           (*Alias)(a),
        AñoPublicacion: a.AñoPublicacion.Format("2006"),
    })
}
