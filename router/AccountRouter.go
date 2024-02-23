package router

import (
	"github.com/gin-gonic/gin"
	"imserver/controller"
)

func AccountRouter(app *gin.Engine) {
	account := app.Group("/api/account")
	{
		account.GET("getCheckCode/:mobile", controller.GetCheckCode)
		account.POST("loginSms", controller.LoginSms)
		account.GET("currentUser", controller.CurrenUser)
		account.GET("accountInfo/:accountId", controller.AccountInfo)
		account.GET("contacts", controller.Contacts)
	}
}
