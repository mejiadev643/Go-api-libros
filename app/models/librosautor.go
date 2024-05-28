package models

type LibrosAutor struct {
    Id      int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
    ISBN   string `gorm:"column:isbn" json:"isbn"`
    AutorId uint   `gorm:"column:id_autor;foreignKey:AutorId" json:"id_autor"`
    LibroId string   `gorm:"column:libro_id;foreignKey:LibroId" json:"libro_id"`
    Libro   Libro  `gorm:"foreignKey:LibroId"`
    Autor   Autor  `gorm:"foreignKey:AutorId"`
}

func (LibrosAutor) TableName() string {
    return "libros_autor"
}