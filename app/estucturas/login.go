package estucturas

type Login struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func(Login) TableName() string {
	return "users"
}