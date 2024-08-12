package model

import "time"

type Raffle struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Prize       float32   `json:"prize"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RaffleRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Prize       float32 `json:"prize"`
}
