package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TransactionHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Transaction history"})
}
