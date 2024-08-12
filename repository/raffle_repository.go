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

func (r *RaffleRepository) GetRaffleById(raffleId int) (*model.Raffle, error) {
	var raffle model.Raffle

	query, err := r.connection.Prepare("SELECT id, name, description, prize FROM raffle WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = query.QueryRow(raffleId).Scan(
		&raffle.Id,
		&raffle.Name,
		&raffle.Description,
		&raffle.Prize,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &raffle, nil
}

func (r *RaffleRepository) CreateRaffle(raffle model.RaffleRequest) (*int, error) {
	var id int

	query, err := r.connection.Prepare("INSERT INTO raffle (name, description, prize) VALUES ($1, $2, $3) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = query.QueryRow(raffle.Name, raffle.Description, raffle.Prize).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &id, nil
}
