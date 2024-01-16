package requests

import (
	"InfotecsTestCase/database"
	"InfotecsTestCase/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateWallet(c *gin.Context) {
	var wallet models.Wallet
	id := uuid.New()
	wallet.Id = id.String()
	wallet.Balance = 100.0

	err := database.CreateWalletDB(&wallet)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Error while creating wallet: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": wallet.Id, "balance": wallet.Balance})
}
