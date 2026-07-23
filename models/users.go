package models

import "time"

type User struct {
	ID        int      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	Birthdate time.Time `json:"birthdate"`
	Gender		string    `json:"gender"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}