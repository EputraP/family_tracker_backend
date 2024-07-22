package handler

import (
	"errors"
	"smart-kost-backend/dto"
	"smart-kost-backend/errs"
	"smart-kost-backend/service"
	"smart-kost-backend/util/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService service.UserService
}

type UserHandlerConfig struct {
	UserService service.UserService
}

func NewUserRawHandler(config UserHandlerConfig) *UserHandler {
	return &UserHandler{
		userService: config.UserService,
	}
}

func (h UserHandler) Login(c *gin.Context) {
	var loginUser dto.LoginBody

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.userService.Login(loginUser)

	if err != nil {
		if errors.Is(err, errs.PasswordDoesntMatch) ||
			errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, 401, errs.UsernamePasswordIncorrect.Error())
			return
		}

		response.UnknownError(c, err)
		return
	}

	response.JSON(c, 200, "Login success", resp)

}

func (h UserHandler) SignUp(c *gin.Context) {
	var createUser dto.User

	if err := c.ShouldBindJSON(&createUser); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.userService.SignUp(createUser)

	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create User Success", resp)
}
