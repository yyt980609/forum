package models

// RegisterForm 注册表单
type RegisterForm struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

// LoginForm 登陆表单
type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// User 用户信息
type User struct {
	UserID       int64  `json:"userId,string" db:"user_id"`
	Username     string `json:"username" db:"username"`
	Password     string `json:"password" db:"password"`
	AccessToken  string `gorm:"-"`
	RefreshToken string `gorm:"-"`
}
