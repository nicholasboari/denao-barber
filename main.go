package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nicholasboari/denao-barber/configs"
	"github.com/nicholasboari/denao-barber/controller"
	"github.com/nicholasboari/denao-barber/helper"
	"github.com/nicholasboari/denao-barber/model"
	"github.com/nicholasboari/denao-barber/repository"
	"github.com/nicholasboari/denao-barber/router"
	"github.com/nicholasboari/denao-barber/service"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")

	// database
	db := configs.DatabaseConnector()
	validate := validator.New()

	db.Table("haircuts").AutoMigrate(&model.Haircut{})

	haircutRepository := repository.NewHaircutRepositoryImpl(db)
	haircutService := service.NewHaircutServiceImpl(haircutRepository, validate)
	haircutController := controller.NewHaircutController(haircutService)

	routes := router.NewRouter(haircutController)

	server := http.Server{
		Addr:    ":8000",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
