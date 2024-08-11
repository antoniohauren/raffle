package repository

import (
	"database/sql"
	"fmt"

	"github.com/antoniohauren/raffle/model"
)

type RaffleRepository struct {
	connection *sql.DB
}

func NewRaffleRepository(connection *sql.DB) RaffleRepository {
	return RaffleRepository{
		connection: connection,
	}
}

func (r *RaffleRepository) GetRaffleList() ([]model.Raffle, error) {
	query := "SELECT id, name, description, prize FROM raffle;"

	rows, err := r.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Raffle{}, err
	}

	var raffleList []model.Raffle
	var raffleEntry model.Raffle

	for rows.Next() {
		err = rows.Scan(
			&raffleEntry.Id,
			&raffleEntry.Name,
			&raffleEntry.Description,
			&raffleEntry.Prize,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Raffle{}, err
		}

		raffleList = append(raffleList, raffleEntry)
	}

	rows.Close()

	return raffleList, nil
}
