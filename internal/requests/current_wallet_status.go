package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurrentWalletStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Current wallet status: "})
}
