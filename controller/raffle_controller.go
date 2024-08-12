package controller

import (
	"encoding/json"
	"net/http"

	"github.com/antoniohauren/raffle/model"
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

func (c *RaffleController) GetRaffleById(w http.ResponseWriter, r *http.Request) {
	id, err := GetIntParam(r, "raffleId")

	if err != nil {
		WriteJson(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := c.raffleUsecase.GetRaffleById(*id)

	if data == nil {
		WriteJson(w, http.StatusNotFound, "Raffle Not Found")
		return
	}

	if err != nil {
		WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	WriteJson(w, http.StatusOK, data)
}

func (c *RaffleController) CreateRaffle(w http.ResponseWriter, r *http.Request) {
	var raffle model.RaffleRequest

	err := json.NewDecoder(r.Body).Decode(&raffle)

	if err != nil {
		WriteJson(w, http.StatusBadRequest, "Invalid body")
		return
	}

	id, err := c.raffleUsecase.CreateRaffle(raffle)

	if err != nil {
		WriteJson(w, http.StatusInternalServerError, "")
	}

	WriteJson(w, http.StatusCreated, id)
}
