package models

import (
	"encoding/json"
	"time"
)

type Autor struct {
	Id   int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"column:name" json:"nombre" validate:"required"`
	Nacionalidad string `gorm:"column:nacionalidad" json:"nacionalidad" validate:"required"`
	AñoNacimiento time.Time `gorm:"column:año_nacimiento;type:date" json:"año-nacimiento" validate:"required"`
	Deleted bool `gorm:"column:deleted;default:false" json:"-" `
}

func (Autor) TableName() string {
	return "autores"
}
func (a *Autor) UnmarshalJSON(data []byte) error {
    type Alias Autor
    aux := &struct {
        AñoNacimiento string `json:"año-nacimiento"`
        *Alias
    }{
        Alias: (*Alias)(a),
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    t, err := time.Parse("2006", aux.AñoNacimiento)
    if err != nil {
        return err
    }

    a.AñoNacimiento = t

    return nil
}

func (a *Autor) MarshalJSON() ([]byte, error) {
    type Alias Autor
    return json.Marshal(&struct {
        *Alias
        AñoNacimiento string `json:"año-nacimiento"`
    }{
        Alias:           (*Alias)(a),
        AñoNacimiento: a.AñoNacimiento.Format("2006"),
    })
}