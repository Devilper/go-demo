package forms

type SavaUserRequest struct {
	UserName string `from:"user_name" json:"user_name" binding:"required"`
	PhoneNum string `from:"phone_num" json:"phone_num" binding:"required,mobile"`
	PassWord string `from:"password" json:"password" binding:"required,min=6,max=10"`
}

type UserLoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
