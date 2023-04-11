package services_test

import (
	"errors"
	"reflect"
	"testing"
	"todo-backend/repositories"
	"todo-backend/services"
)

var userRepoMock = repositories.NewUserRepositoryMock()
var userService = services.NewUserService(userRepoMock)

func TestLoginShouldBePass(t *testing.T) {

	loginTest := []struct {
		name string
		user services.UserLoginRequest
		want *services.UserLoginResponse
	}{
		{
			name: "jonh logined",
			user: services.UserLoginRequest{Username: "jonh", Password: "jonhpass"},
			want: &services.UserLoginResponse{Username: "jonh", Token: ""},
		},
		{
			name: "bob logined",
			user: services.UserLoginRequest{Username: "bob", Password: "strongpassword"},
			want: &services.UserLoginResponse{Username: "bob", Token: ""},
		},
		{
			name: "alan logined",
			user: services.UserLoginRequest{Username: "alan", Password: "alanpass"},
			want: &services.UserLoginResponse{Username: "alan", Token: ""},
		},
		{
			name: "bas logined",
			user: services.UserLoginRequest{Username: "bas", Password: "baspass"},
			want: &services.UserLoginResponse{Username: "bas", Token: ""},
		},
	}

	for _, tt := range loginTest {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userService.Login(tt.user)
			if err != nil {
				t.Errorf(err.Error())
			}
			if got.Username != tt.want.Username || got.Token == "" {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestLoginShouldBeFail(t *testing.T) {

	loginTest := []struct {
		name string
		user services.UserLoginRequest
		want error
	}{
		{
			name: "jonh is login with incorrect password",
			user: services.UserLoginRequest{Username: "jonh", Password: "incorrectpass"},
			want: errors.New("username or password is incorrect"),
		},
		{
			name: "jonh is login with incorret username",
			user: services.UserLoginRequest{Username: "incorrectusername", Password: "incorrectpass"},
			want: errors.New("username or password is incorrect"),
		},
		{
			name: "bob is loging with empty username",
			user: services.UserLoginRequest{Username: "", Password: "strongpassword"},
			want: errors.New("username or password is empty"),
		},
		{
			name: "alan is loging with empty pass",
			user: services.UserLoginRequest{Username: "alan", Password: ""},
			want: errors.New("username or password is empty"),
		},
	}

	for _, tt := range loginTest {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userService.Login(tt.user)
			if err == nil {
				t.Errorf("don't have a error but i want")
				return
			}
			if err.Error() != tt.want.Error() {
				t.Errorf("got %v want %v ", err, tt.want.Error())
			}
		})
	}
}

func TestSignUpShouldBePass(t *testing.T) {

	signUpTest := []struct {
		name string
		user services.UserSignUpRequest
		want *services.UserResponse
	}{
		{
			name: "new user 1",
			user: services.UserSignUpRequest{Username: "newuser1", Password: "strongpass"},
			want: &services.UserResponse{Username: "newuser1"},
		},
		{
			name: "new user 2",
			user: services.UserSignUpRequest{Username: "newuser2", Password: "strongpass"},
			want: &services.UserResponse{Username: "newuser2"},
		},
		{
			name: "new user 3",
			user: services.UserSignUpRequest{Username: "newuser3", Password: "strongpass"},
			want: &services.UserResponse{Username: "newuser3"},
		},
	}

	for _, tt := range signUpTest {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userService.SignUp(tt.user)
			if err != nil {
				t.Errorf(err.Error())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestSignUpShouldBeFail(t *testing.T) {

	signUpTest := []struct {
		name string
		user services.UserSignUpRequest
		want error
	}{
		{
			name: "singup with empty username",
			user: services.UserSignUpRequest{Username: "", Password: "password"},
			want: errors.New("username or password is empty"),
		},
		{
			name: "singup with empty password",
			user: services.UserSignUpRequest{Username: "username", Password: ""},
			want: errors.New("username or password is empty"),
		},
		{
			name: "singup with existing username",
			user: services.UserSignUpRequest{Username: "jonh", Password: "password"},
			want: errors.New("username is duplicate"),
		},
	}

	for _, tt := range signUpTest {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userService.SignUp(tt.user)
			if err == nil {
				t.Errorf("don't have a error but i want")
				return
			}
			if err.Error() != tt.want.Error() {
				t.Errorf("got %v want %v ", err, tt.want.Error())
			}
		})
	}
}
