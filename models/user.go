package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string    `json:"password"`
	Nickname  string    `json:"nickname"`
}



