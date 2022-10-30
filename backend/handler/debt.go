package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-shinzato/pay-settle/model"
	"net/http"
)

// GetDebts responds with the list of all albums as JSON.
func GetDebts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.GetDebts())
}

// Register handles a requests for registering one debt
func Register(c *gin.Context) {
	// parse request body
	var debt model.Debt
	c.BindJSON(&debt)

	if debt.Price <= 0 {
		c.String(http.StatusBadRequest, fmt.Sprintf("Price should be positive integer, but got %v", debt.Price))
		return
	}

	if !((debt.Debtor == "miryu" && debt.Lender == "pon") || (debt.Debtor == "pon" && debt.Lender == "miryu")) {
		c.String(http.StatusBadRequest, fmt.Sprintf("(debtor, lender) should be (miryu, pon) or (pon, miryu), but (debtor, lender) is (%v, %v)", debt.Debtor, debt.Lender))
		return
	}

	if err := model.Register(&debt); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "Success")
}

// Calculate responds with the price someone have to pay.
func Calculate(c *gin.Context) {
	debts := model.GetDebts()
	sum := 0
	for _, debt := range debts {
		if debt.Completed == 1 {
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
