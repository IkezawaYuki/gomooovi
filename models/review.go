package models

import "time"

type Review struct {
	Id        int       `json:"id"`
	Rate      int       `json:"rate"`
	Comment   string    `json:"comment"`
	ProductId int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    int       `json:"user_id"`
}

type Reviews struct {
	Reviews   []Review 	`json:"review"`
}