package user

type UserCreateRequest struct {
	Fullname       string
	Email          string
	Password       string
	ForgotPassword string
	RoleId         int
}
