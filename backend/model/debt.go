package model

import (
	"fmt"
	"time"
)

// Debt represents data about a record debt.
type Debt struct {
	ID        string    `json:"id"`
	Debtor    string    `json:"debtor"`
	Lender    string    `json:"lender"`
	Memo      string    `json:"memo"`
	Price     int       `json:"price"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Register registers a debt
func Register(debt *Debt) error {
	// TODO insert debt to DataBase

	// print debt for just debugging
	fmt.Println("inserted: debt.Price = ", debt.Price)
	fmt.Println("inserted: debt.Lender = ", debt.Lender)
	fmt.Println("inserted: debt.Debtor = ", debt.Debtor)

	return nil
}

// GetDebts returns slice to seed record debt data.
func GetDebts() []Debt {
	// TODO DB から借金データを取ってくる
	return []Debt{
		{
			ID:        "1",
			Debtor:    "pon",
			Lender:    "miryu",
			Memo:      "葡庵、ランチ",
			Price:     2080,
			Completed: false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			Debtor:    "pon",
			Lender:    "miryu",
			Memo:      "一慶",
			Price:     7500,
			Completed: false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "3",
			Debtor:    "miryu",
			Lender:    "pon",
			Memo:      "nosh",
			Price:     800,
			Completed: false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
