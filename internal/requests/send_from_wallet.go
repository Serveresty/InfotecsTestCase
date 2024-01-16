package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMoneyFromWallet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Send money from wallet"})
}
