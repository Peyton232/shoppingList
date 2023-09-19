package models

import (
	"strconv"
)

type Item struct {
	ID          string `json:"id" pg:"id,pk"`
	Name        string `json:"name" pg:"name"`
	Description string `json:"description" pg:"description"`
	Amount      string `json:"amount" pg:"amount"`
	Completed   bool   `json:"completed" pg:"completed"`
	CreatedAt   string `json:"created_at" pg:"created_at"`
	UpdatedAt   string `json:"updated_at" pg:"updated_at"`
}

func ParseItemID(itemIDStr string) (int, error) {
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		return 0, err
	}
	return itemID, nil
}
