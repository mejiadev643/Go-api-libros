package models

type Editorial struct {
	Id   int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"column:name" json:"nombre" validate:"required"`
	Direccion string `gorm:"column:direccion" json:"direccion" validate:"required"`
	Contacto string `gorm:"column:contacto" json:"contacto" validate:"required"`
	Deleted bool `gorm:"column:deleted;default:false" json:"-" `
}

func (Editorial) TableName() string {
	return "editoriales"
}