package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shinzato/pay-settle/model"
	"net/http"
)

// GetDebts responds with the list of all albums as JSON.
func GetDebts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.GetDebts())
}

// Calculate responds with the price someone have to pay.
func Calculate(c *gin.Context) {
	debts := model.GetDebts()
	sum := 0
	for _, debt := range debts {
		if debt.Completed {
			continue
		}
		if debt.Debtor == "miryu" {
			sum += debt.Price
		} else {
			sum -= debt.Price
		}
	}

	type Response struct {
		Debtor string `json:"debtor"`
		Lender string `json:"lender"`
		Price  int    `json:"price"`
	}

	var res Response

	if sum > 0 {
		res = Response{
			Debtor: "miryu",
			Lender: "pon",
			Price:  sum / 2,
		}
	} else {
		res = Response{
			Debtor: "pon",
			Lender: "miryu",
			Price:  -sum / 2,
		}
	}

	c.IndentedJSON(http.StatusOK, res)
}

func Resister(c *gin.Context) {
	
	// TODO 登録処理
	// request body を c から受け取る
	// request body のバリデーションチェック

	// チェックが OK ならば DB へ登録
	if err := model.Register(debt); err != nil {
		// do error handling
		// do 500 bann wo kaesu
	}
	c.IndentedJSON(http.StatusOK)
}
