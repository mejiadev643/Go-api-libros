package models

type Description struct {
	Id               int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Id_user          uint    `gorm:"column:code;foreignKey:Id_user" json:"usuario_id"`
	First_name       string `gorm:"column:nombres" json:"nombres"`
	Last_name        string `gorm:"column:apellidos" json:"apellidos"`
	Direction        string `gorm:"column:direccion" json:"direccion"`
	Telephone_number string `gorm:"column:telefono" json:"telefono"`
}

// TableName sobrescribe el nombre de la tabla usado por Usuario para `mnt_usuario`.
func (Description) TableName() string {
	return "description"
}
