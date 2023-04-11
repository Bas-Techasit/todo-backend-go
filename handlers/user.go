package handlers

import (
	"todo-backend/services"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) Login(c *fiber.Ctx) error {
	userReq := services.UserLoginRequest{}
	if err := c.BodyParser(&userReq); err != nil {
		return err
	}
	userRes, err := h.userSrv.Login(userReq)
	if err != nil {
		return err
	}

	return c.JSON(userRes)
}

func (h userHandler) SignUp(c *fiber.Ctx) error {
	user := services.UserSignUpRequest{}
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	res, err := h.userSrv.SignUp(user)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}
