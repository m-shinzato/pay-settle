package model

import (
	"time"
    "database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

// Debt represents data about a record debt.
type Debt struct {
	ID        string    `json:"id"`
	Debtor    string    `json:"debtor"`
	Lender    string    `json:"lender"`
	Memo      string    `json:"memo"`
	Price     int       `json:"price"`
	Completed int       `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Register registers a debt
func Register(debt *Debt) error {
	// DBを開く  なければ作成される
    DbConnection, _ := sql.Open("sqlite3", "../pay-settle.sqlite")
    defer DbConnection.Close()

    sql := `insert into debts (lender, debtor, price, memo) values (?, ?, ?, ?);`

    // 実行 結果は返ってこない為、_にする
    if _, err := DbConnection.Exec(sql, debt.Lender, debt.Debtor, debt.Price, debt.Memo); err != nil {
        return err
    }

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
			Completed: 0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			Debtor:    "pon",
			Lender:    "miryu",
			Memo:      "一慶",
			Price:     7500,
			Completed: 0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "3",
			Debtor:    "miryu",
			Lender:    "pon",
			Memo:      "nosh",
			Price:     800,
			Completed: 0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
