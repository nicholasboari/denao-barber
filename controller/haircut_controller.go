package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nicholasboari/denao-barber/data/request"
	"github.com/nicholasboari/denao-barber/data/response"
	"github.com/nicholasboari/denao-barber/model"
	"github.com/nicholasboari/denao-barber/service"
	"github.com/rs/zerolog/log"
)

type HaircutController struct {
	haircutService service.HaircutService
}

func NewHaircutController(service service.HaircutService) *HaircutController {
	return &HaircutController{
		haircutService: service,
	}
}

func (controller *HaircutController) Create(ctx *gin.Context) {
	log.Info().Msg("create haircut")
	var request request.CreateHaircutRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	// yyyy-MM-dd
	// 2006-01-13 15:00:00
	date, err := time.Parse(time.DateTime, request.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	haircut := model.Haircut{NameClient: request.NameClient, Price: request.Price, Date: date}
	controller.haircutService.Create(&haircut)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{"data": haircut})
}

func (controller *HaircutController) GetHaircutByID(ctx *gin.Context) {
	log.Info().Msg("get haircut by ID")
	var idString = ctx.Param("id")
	ID, err := uuid.Parse(idString)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid UUID"})
		return
	}
	var data *model.Haircut
	data, err = controller.haircutService.GetHaircutByID(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status: "error to get haircut",
			Data:   err,
		})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (controller *HaircutController) GetAllHaircuts(ctx *gin.Context) {
	log.Info().Msg("get all haircuts")
	var data []*model.Haircut
	data, err := controller.haircutService.GetAllHaircuts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status: "error to get all haircuts",
			Data:   err,
		})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}
