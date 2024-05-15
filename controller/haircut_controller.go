package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicholasboari/denao-barber/data/request"
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
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	// yyyy-MM-dd
	// 2006-01-13 15:00:00
	date, err := time.Parse(time.DateTime, request.Date)
	if err != nil {
		ctx.JSON(404, err)
		return
	}
	haircut := model.Haircut{NameClient: request.NameClient, Price: request.Price, Date: date}
	controller.haircutService.Create(&haircut)
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{"data": haircut})
}
