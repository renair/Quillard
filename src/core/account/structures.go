package account

//DB Objects

import (
	"core/positions"
)

//Type represents `account` table row
type Account struct {
	Id             int64
	Email          string
	Password       string
	rawPassword    string
	HomePositionId int64
}

//JSON Request Objects

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Position positions.Position `json:"home_position"`
}
