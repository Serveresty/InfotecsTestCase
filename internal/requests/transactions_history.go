package requests

import (
	"InfotecsTestCase/cerr"
	"InfotecsTestCase/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TransactionHistory(c *gin.Context) {
	id := c.Param("id")
	transactions, err := database.GetTransactionHistory(id)
	if err != nil {
		if err == cerr.WalletNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": 404, "error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "data": transactions})
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "data": transactions})
}
