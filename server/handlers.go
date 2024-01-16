package server

import (
	"InfotecsTestCase/internal/requests"

	"github.com/gin-gonic/gin"
)

func AllRequests(router *gin.Engine) {
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.POST("wallet", requests.CreateWallet)

			wallet := v1.Group("wallet")
			{
				wallet.POST(":id", requests.CurrentWalletStatus)

				id := wallet.Group(":id")
				{
					id.POST("send", requests.SendMoneyFromWallet)
					id.POST("history", requests.TransactionHistory)
				}
			}
		}
	}
}
