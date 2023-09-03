package models

// RegisterForm 注册表单
type RegisterForm struct {
	Username        string `json:"username" binding:"required" zh:"用户名"`
	Password        string `json:"password" binding:"required" zh:"密码"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password" zh:"确认密码"`
}

// LoginForm 登陆表单
type LoginForm struct {
	Username string `json:"username" binding:"required" zh:"用户名"`
	Password string `json:"password" binding:"required" zh:"密码"`
}

// User 用户信息
type User struct {
	UserID       int64  `json:"userId,string" db:"user_id"`
	Username     string `json:"username" db:"username"`
	Password     string `json:"password" db:"password"`
	AccessToken  string `json:"accessToken" gorm:"-"`
	RefreshToken string `json:"refreshToken" gorm:"-"`
}
