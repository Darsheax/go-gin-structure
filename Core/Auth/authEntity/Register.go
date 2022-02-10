package authEntity

type Register struct {
	UserName string `form:"username" json:"username" binding:"alpha,required"`
	Email    string `form:"email" json:"email" binding:"email,required"`
	Password string `form:"password" json:"password" binding:"required"`
}
