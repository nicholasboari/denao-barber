package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	err = controller.userService.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status: "error to create user",
			Data:   err,
		})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func (controller *UserController) GetUserByID(ctx *gin.Context) {
	log.Info().Msg("get user by ID")
	var idString = ctx.Param("id")
	ID, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid UUID"})
		return
	}
	var data *model.User
	data, err = controller.userService.GetUserByID(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status: "error to get user",
			Data:   err,
		})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{"data": data})
}

func (controller *UserController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete user by ID")
	var idString = ctx.Param("id")
	ID, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid UUID"})
		return
	}
	err = controller.userService.Delete(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status: "error to delete user",
			Data:   err.Error(),
		})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusNoContent, nil)
}

func (controller *UserController) Update(ctx *gin.Context) {
	log.Info().Msg("update user")
	var request request.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ID, err := uuid.Parse(request.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	user := &model.User{
		ID:   ID,
		Name: request.Name,
	}
	userUpdated, err := controller.userService.Update(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status: "error to update user",
			Data:   err,
		})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"data": userUpdated})
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
