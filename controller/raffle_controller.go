package controller

import (
	"net/http"

	"github.com/antoniohauren/raffle/usecase"
)

type RaffleController struct {
	raffleUsecase usecase.RaffleUsecase
}

func NewRaffleController(usecase usecase.RaffleUsecase) RaffleController {
	return RaffleController{
		raffleUsecase: usecase,
	}
}

func (c *RaffleController) GetRaffleList(w http.ResponseWriter, r *http.Request) {
	data, err := c.raffleUsecase.GetRaffleList()

	if err != nil {
		WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	WriteJson(w, http.StatusOK, data)
}
