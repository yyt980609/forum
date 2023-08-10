package models

type RegisterForm struct {
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	UserID       uint64 `json:"userID,string" db:"user_id"`
	Username     string `json:"username" db:"username"`
	Password     string `json:"password" db:"password"`
	AccessToken  string `gorm:"-"`
	RefreshToken string `gorm:"-"`
}
