package models

type LibrosAutor struct {
    Id      int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    ISBN   string `gorm:"column:isbn" json:"isbn" validate:"required"`
    AutorId uint   `gorm:"column:id_autor;foreignKey:AutorId" json:"id_autor" validate:"required"`
    LibroId string   `gorm:"column:libro_id;foreignKey:LibroId" json:"libro_id" validate:"required"`
	Deleted bool `gorm:"column:deleted;default:false" json:"-" `
    Libro   Libro  `gorm:"foreignKey:LibroId" `
    Autor   Autor  `gorm:"foreignKey:AutorId" `
}

type LibrosAutores struct {
    Id      int    `json:"id"`
    ISBN   string `json:"isbn"`
    LibroId string   `json:"libro_id"`
    Libro   Libro  `json:"libro"`
    Autores   []Autor `json:"autores"`
}

func (LibrosAutor) TableName() string {
    return "libros_autor"
}