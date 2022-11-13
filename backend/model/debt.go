package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

// Debt represents data about a record debt.
type Debt struct {
	ID        string `json:"id"`
	Debtor    string `json:"debtor"`
	Lender    string `json:"lender"`
	Price     int    `json:"price"`
	Memo      string `json:"memo"`
	Completed int    `json:"completed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Pay registers a debt
func Pay(debt *Debt) error {
	DbConnection, _ := sql.Open("sqlite3", "../pay-settle.sqlite")
	defer DbConnection.Close()

	sql := `insert into debts (lender, debtor, price, memo) values (?, ?, ?, ?);`
	_, err := DbConnection.Exec(sql, debt.Lender, debt.Debtor, debt.Price, debt.Memo)
	return err
}

// GetDebts returns slice to seed record debt data.
func GetDebts() ([]*Debt, error) {
	DbConnection, _ := sql.Open("sqlite3", "../pay-settle.sqlite")
	defer DbConnection.Close()

	sql := `select * from debts where completed = 0`
	recodes, err := DbConnection.Query(sql)
	if err != nil {
		return nil, err
	}

	debts := make([]*Debt, 0)
	for recodes.Next() {
		var debt Debt
		err := recodes.Scan(&debt.ID, &debt.Debtor, &debt.Lender, &debt.Price, &debt.Memo, &debt.Completed, &debt.CreatedAt, &debt.UpdatedAt)
		if err != nil {
			return nil, err
		}
		debts = append(debts, &debt)
	}
	return debts, nil
}

// DeleteDebts debts
func DeleteDebts(debt *Debt) error {
	DbConnection, _ := sql.Open("sqlite3", "../pay-settle.sqlite")
	defer DbConnection.Close()

	sql := `delete from debts where id = ?`
	if _, err := DbConnection.Exec(sql, debt.ID); err != nil {
		return err
	}

	return nil
}

// Settle a debt
func Settle() error {
	DbConnection, _ := sql.Open("sqlite3", "../pay-settle.sqlite")
	defer DbConnection.Close()

	sql := `update debts set completed = 1 where completed = 0;`

	_, err := DbConnection.Exec(sql)
	return err
}
