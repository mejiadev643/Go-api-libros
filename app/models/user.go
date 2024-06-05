package models

type User struct {
	Id             int           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Code           int           `gorm:"column:code" json:"code" validate:"required"`
	Email          string        `gorm:"column:email" json:"email" validate:"required,email"`
	Nombres        string        `gorm:"column:nombres" json:"nombre" validate:"required"`
	Password       string        `gorm:"column:password" json:"password" validate:"required"`
	Super_password string        `gorm:"column:super_password" json:"super-password" validate:"required"`
	Last_conection string        `gorm:"column:last_connection" json:"last_connection"`
	Deleted        bool          `gorm:"column:deleted;default:false" json:"-"`
	Descriptions   Description `gorm:"foreignKey:Id_user"`
}
type UserStruct struct {
	Id             int           `json:"id"`
	Code           int           `json:"code"`
	Email          string        `json:"email"`
	Nombres        string        `json:"nombre"`
	Last_conection string        `json:"last_connection"`
	Descriptions   Description `gorm:"foreignKey:Id_user"`
}

// TableName sobrescribe el nombre de la tabla usado por Usuario para `mnt_usuario`.
func (User) TableName() string {
	return "users"
}

func (UserStruct) TableName() string {
	return "users"
}
