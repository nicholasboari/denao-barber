package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nicholasboari/denao-barber/controller"
)

func NewRouter(haircutController *controller.HaircutController) *gin.Engine {
	router := gin.Default()

	haircutRouter := router.Group("/haircuts")
	haircutRouter.POST("", haircutController.Create)

	return router
}
