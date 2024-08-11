package router

import (
	"database/sql"
	"net/http"

	"github.com/antoniohauren/raffle/controller"
	"github.com/antoniohauren/raffle/repository"
	"github.com/antoniohauren/raffle/usecase"
)

func NewRaffleRouter(connection *sql.DB) *http.ServeMux {
	repository := repository.NewRaffleRepository(connection)
	usecase := usecase.NewRaffleUsecase(repository)
	controller := controller.NewRaffleController(usecase)

	raffleRouter := http.NewServeMux()

	raffleRouter.HandleFunc("GET /", controller.GetRaffleList)

	return raffleRouter
}
