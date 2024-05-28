package models

type Autor struct {
	Id   int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"column:name" json:"nombre"`
	Nacionalidad string `gorm:"column:nacionalidad" json:"nacionalidad"`
	FechaNacimiento string `gorm:"column:fecha_nacimiento" json:"fecha_nacimiento"`
	LibrosAutor []LibrosAutor `gorm:"foreignKey:AutorId"`
}

func (Autor) TableName() string {
	return "autores"
}