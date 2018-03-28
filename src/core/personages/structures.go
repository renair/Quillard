package personages

import (
	"core/positions"
)

type Personage struct {
	Id         int64
	AccountId  int64
	Name       string
	PositionId int64
}

type PersonageResponse struct {
	Id       int64              `json:"id"`
	Name     string             `json:"name"`
	Position positions.Position `json:"position"`
}

type PersonageRequest struct {
	Name string `json:"name"`
}

//function to convert types
func (pers Personage) toResponse() PersonageResponse {
	return PersonageResponse{
		Id:       pers.Id,
		Name:     pers.Name,
		Position: *positions.GetPositionById(pers.PositionId),
	}
}
