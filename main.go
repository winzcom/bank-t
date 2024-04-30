package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/winzcom/bank-t/db"
	"github.com/winzcom/bank-t/transaction"
)

func main() {
	db := db.GetStorage()
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/debit", func(ctx *gin.Context) {
		defer ctx.Request.Body.Close()
		var requestBody transaction.Transaction

		ctx.Header("Content-Type", "application/json")

		if err := ctx.BindJSON(&requestBody); err != nil {
			// DO SOMETHING WITH THE ERROR
			fmt.Println("whasy iasda ", err)
		}
		err := transaction.Debit(transaction.Transaction{
			Amount:     requestBody.Amount,
			Account_id: "abc",
			Db:         db,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "Failed",
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Success",
		})
	})
	r.POST("/credit", func(ctx *gin.Context) {
		defer ctx.Request.Body.Close()
		var requestBody transaction.Transaction

		ctx.Header("Content-Type", "application/json")

		if err := ctx.BindJSON(&requestBody); err != nil {
			// DO SOMETHING WITH THE ERROR
			fmt.Println("whasy iasda ", err)
		}
		if err := transaction.Credit(transaction.Transaction{
			Amount:     requestBody.Amount,
			Account_id: "abc",
		}); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "Failed",
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Success",
		})
	})
	r.Run()
}
