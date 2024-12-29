package models

type TopUpRequest struct {
	Amount int `json:"amount" validate:"required,min=1"`
}
