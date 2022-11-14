package models

type Auth struct {
	BaseModel
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
