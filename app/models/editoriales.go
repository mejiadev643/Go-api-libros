package models

type Editorial struct {
	Id   int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"column:name" json:"nombre"`
	Direccion string `gorm:"column:direccion" json:"direccion"`
	Contacto string `gorm:"column:contacto" json:"contacto"`
	Libros []Libro `gorm:"foreignKey:EditorialId"`
}

func (Editorial) TableName() string {
	return "editoriales"
}