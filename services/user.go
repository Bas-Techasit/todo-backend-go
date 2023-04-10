package services

type UserService interface {
	Login(UserLoginRequest) (*UserLoginResponse, error)
	SignUp(UserSignUpRequest) (*UserResponse, error)
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
