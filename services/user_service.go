package services

import (
	"todo-backend/logs"
	"todo-backend/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

const SECRET = "techasit"

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) Login(userReq UserLoginRequest) (*UserLoginResponse, error) {
	if userReq.Username == "" || userReq.Password == "" {
		return nil, fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"username or password is empty",
		)
	}

	user, err := s.userRepo.GetUser(userReq.Username)
	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"username or password is incorrect",
		)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		// logs.Error(err)
		return nil, fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"username or password is incorrect",
		)
	}

	// generate token
	cliams := jwt.RegisteredClaims{Issuer: userReq.Username}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)

	var token string
	token, err = jwtToken.SignedString([]byte(viper.GetString("app.secret")))
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	userLonginResponse := UserLoginResponse{
		Username: user.Username,
		Token:    token,
	}

	return &userLonginResponse, nil
}

func (s userService) SignUp(userReq UserSignUpRequest) (*UserResponse, error) {
	if userReq.Username == "" || userReq.Password == "" {
		return nil, fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"username or password is empty",
		)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.MinCost)
	if err != nil {
		logs.Error(err)
		return nil, fiber.NewError(
			fiber.StatusInternalServerError,
			"unexpected error",
		)
	}

	var user *repositories.User
	user, err = s.userRepo.CreateUser(repositories.User{
		Username: userReq.Username,
		Password: string(hashedPass),
	})
	if err != nil {
		return nil, fiber.NewError(
			fiber.ErrUnprocessableEntity.Code,
			"username is duplicate",
		)
	}

	userResponse := UserResponse{
		Username: user.Username,
	}
	return &userResponse, nil
}
