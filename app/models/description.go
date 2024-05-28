package models

type Description struct {
	Id               int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Id_user          uint    `gorm:"column:code;foreignKey:Id_user" json:"usuario_id"`
	First_name       string `gorm:"column:nombres" json:"nombre"`
	Last_name        string `gorm:"column:password" json:"-"`
	Direction        string `gorm:"column:super_password" json:"-"`
	Telephone_number string `gorm:"column:last_connection" json:"last_connection"`
}

// TableName sobrescribe el nombre de la tabla usado por Usuario para `mnt_usuario`.
func (Description) TableName() string {
	return "description"
}
