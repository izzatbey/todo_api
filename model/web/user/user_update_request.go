package user

type UserUpdateRequest struct {
	Fullname       string `validate:"required,min=1,max=100" json:"fullname"`
	Email          string `validate:"required,email,min=1,max=100" json:"email"`
	Password       string `validate:"required,min=1,max=255" json:"password"`
	ForgotPassword string `validate:"min=0,max=255" json:"forgot_password"`
	RoleId         int    `validate:"required,numeric" json:"role_id"`
	Id             int    `validate:"required"`
}
