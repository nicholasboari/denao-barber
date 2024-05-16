package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/nicholasboari/denao-barber/data/request"
	"github.com/nicholasboari/denao-barber/data/response"
	"github.com/nicholasboari/denao-barber/model"
	"github.com/nicholasboari/denao-barber/service"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (controller *UserController) Create(ctx *gin.Context) {
	log.Info().Msg("create user")
	var request request.CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if !valid(request.Email) {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "email is not valid",
		})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	user := model.User{Name: request.Name, Email: request.Email, Password: password}
	controller.userService.Create(&user)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
