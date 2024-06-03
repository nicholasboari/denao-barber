package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nicholasboari/denao-barber/controller"
)

func NewRouter(haircutController *controller.HaircutController, userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	haircutRouter := router.Group("/haircuts")
	userRouter := router.Group("/users")
	haircutRouter.POST("/", haircutController.Create)
	haircutRouter.DELETE("/:id", haircutController.Delete)
	haircutRouter.GET("/", haircutController.GetAllHaircuts)
	haircutRouter.GET("/:id", haircutController.GetHaircutByID)

	userRouter.POST("/", userController.Create)
	userRouter.GET("/:id", userController.GetUserByID)
	userRouter.GET("/", userController.FindAll)
	userRouter.DELETE("/:id", userController.Delete)
	userRouter.PUT("/", userController.Update)
	return router
}
