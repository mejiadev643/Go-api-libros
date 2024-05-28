package models

type Genero struct {
	Id   int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"column:name" json:"nombre" validate:"required"`
	Deleted bool `gorm:"column:deleted;default:false" json:"-" `
}

func (Genero) TableName() string {
	return "generos"
}