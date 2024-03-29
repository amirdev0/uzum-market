package models

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int32     `json:"quantity"`
}

type ShortProductInfo struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Quantity int32     `json:"quantity"`
}
