package user

type UserUpdateRequest struct {
	Fullname       string
	Email          string
	Password       string
	ForgotPassword string
	RoleId         int
	Id             int
}
