package models

type PayRequest struct {
	Amount int `json:"amount" validate:"required,min=1"`
}
