package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	apiPrefix := "/api"
	g := app.Group(apiPrefix)
	//setting
	RegisterRouterSetting(g)
}
