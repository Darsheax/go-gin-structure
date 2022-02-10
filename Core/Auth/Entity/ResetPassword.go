package authEntity

type ResetPassword struct {
	Password       string `form:"password" json:"password" binding:"required"`
	PasswordVerify string `form:"password_verify" json:"password_verify" binding:"required"`
}
