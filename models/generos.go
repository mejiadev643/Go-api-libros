package models

type Genero struct {
	Id   int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"column:name" json:"nombre"`
	Libros []Libro `gorm:"foreignKey:GeneroId"`
}

func (Genero) TableName() string {
	return "generos"
}