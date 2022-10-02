package model

import "time"

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
