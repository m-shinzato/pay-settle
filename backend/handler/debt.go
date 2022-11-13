package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-shinzato/pay-settle/model"
)

// GetDebts responds with the list of all albums as JSON.
func GetDebts(c *gin.Context) {
	debts, err := model.GetDebts()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, debts)
}

// Pay handles a request for registering one debt.
func Pay(c *gin.Context) {
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

	if err := model.Pay(&debt); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "Success")
}

// Calculate responds with the price someone have to pay.
func Calculate(c *gin.Context) {
	debts, err := model.GetDebts()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
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

// DeleteDebts handles a request for deleting debts.
func DeleteDebts(c *gin.Context) {
	// parse request body
	var debt model.Debt
	c.BindJSON(&debt)
	if debt.ID == "" {
		c.String(http.StatusBadRequest, "ID should not be empty string")
		return
	}

	if err := model.DeleteDebts(&debt); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "Success")
}

// Settle handles a request for settling debts.
func Settle(c *gin.Context) {
	err := model.Settle()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, "Success")
}
