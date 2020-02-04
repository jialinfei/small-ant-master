package routers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	apiPrefix := "/api"
	g := app.Group(apiPrefix)
	//setting
	RegisterRouterSetting(g)
}
