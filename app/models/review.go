package models

import "time"

type Review struct {
	ID        int       `json:"id"`
	Rate      int       `json:"rate"`
	Comment   string    `json:"comment"`
	ProductID int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `json:"user_id"`
}
