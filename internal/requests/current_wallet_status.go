package requests

import (
	"InfotecsTestCase/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurrentWalletStatus(c *gin.Context) {
	id := c.Param("id")
	wallet, err := database.GetWallet(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet is not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": wallet.Id, "balance": wallet.Balance})
}
