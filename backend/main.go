package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shinzato/pay-settle/handler"
)

func main() {
	router := gin.Default()
	router.GET("/debts", handler.GetDebts)
	router.GET("/calculate", handler.Calculate)
	router.POST("/register", handler.Register)

	router.Run("localhost:8080")
}
