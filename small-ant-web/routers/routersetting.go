package routers

import (
	"github.com/gin-gonic/gin"
	"small-ant-parent/small-ant-web/controller"
)

func RegisterRouterSetting(app *gin.RouterGroup) {
	user := controller.User{}
	app.GET("/user/get", user.Info)
	app.POST("/user/login", user.Login)
}
