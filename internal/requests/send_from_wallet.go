package requests

import (
	"InfotecsTestCase/cerr"
	"InfotecsTestCase/database"
	"InfotecsTestCase/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMoneyFromWallet(c *gin.Context) {
	from_id := c.Param("id")

	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "error": "Wrong input data"})
		return
	}
	transaction.From = from_id

	err := database.SendMoney(&transaction)
	if err != nil {
		switch err {
		case cerr.SenderWalletNotFound:
			c.JSON(http.StatusNotFound, gin.H{"status": 404, "error": err.Error()})
			return
		case cerr.ReceiverWalletNotFound, cerr.LessMoney:
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "Transaction success"})
}
