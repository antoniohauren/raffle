package usecase

import (
	"github.com/antoniohauren/raffle/model"
	"github.com/antoniohauren/raffle/repository"
)

type RaffleUsecase struct {
	raffleRepository repository.RaffleRepository
}

func NewRaffleUsecase(repository repository.RaffleRepository) RaffleUsecase {
	return RaffleUsecase{
		raffleRepository: repository,
	}
}

func (u *RaffleUsecase) GetRaffleList() ([]model.Raffle, error) {
	return u.raffleRepository.GetRaffleList()
}

func (u *RaffleUsecase) GetRaffleById(raffleId int) (*model.Raffle, error) {
	return u.raffleRepository.GetRaffleById(raffleId)
}
