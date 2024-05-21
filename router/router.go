package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nicholasboari/denao-barber/controller"
)

func NewRouter(haircutController *controller.HaircutController, userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	haircutRouter := router.Group("/haircuts")
	userRouter := router.Group("/users")
	haircutRouter.POST("", haircutController.Create)
	userRouter.POST("", userController.Create)
	userRouter.GET("/:id", userController.GetUserByID)
	return router
}
