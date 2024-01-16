package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWallet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Created wallet"})
}
