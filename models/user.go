package models

type User struct {
	Id             int           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Code           int           `gorm:"column:code"`
	Nombres        string        `gorm:"column:nombres" json:"nombre"`
	Password       string        `gorm:"column:password" json:"-"`
	Super_password string        `gorm:"column:super_password" json:"-"`
	Last_conection string        `gorm:"column:last_connection" json:"last_connection"`
	Descriptions   []Description `gorm:"foreignKey:Id_user"`
}

// TableName sobrescribe el nombre de la tabla usado por Usuario para `mnt_usuario`.
func (User) TableName() string {
	return "users"
}
