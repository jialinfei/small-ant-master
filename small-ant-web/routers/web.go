package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/it234/goapp/pkg/convert"
	"github.com/it234/goapp/pkg/logger"
	"net/http"
	"os"
	"small-ant-parent/small-ant-commn/config"
	"small-ant-parent/small-ant-commn/middleware"
	"time"
)

// 运行
func Run(configPath string) {
	if configPath == "" {
		configPath = "./config.yaml"
	}
	// 加载配置
	config, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	logger.InitLog("debug", "./log/logb.log")
	initDB(config)
	initWeb(config)
	logger.Debug(config.Web.Domain + "站点已启动...")
}

func initDB(configBase *config.Config) {
	config.InitDB(configBase)
}

func initWeb(config *config.Config) {
	gin.SetMode(gin.DebugMode) //调试模式
	app := gin.New()
	app.NoRoute(middleware.NoRouteHandler())
	// 崩溃恢复
	path, _ := os.Getwd()
	app.Use(middleware.RecoveryMiddleware())
	app.LoadHTMLGlob(fmt.Sprintf("%s%s%s", path, config.Web.StaticPath, "view/*.html"))
	app.Static("/static", fmt.Sprintf("%s%s%s", path, config.Web.StaticPath, "view/static"))
	app.Static("/resource", fmt.Sprintf("%s%s%s", path, config.Web.StaticPath, "view/resource"))
	RegisterRouter(app)
	go initHTTPServer(config, app)
}

// InitHTTPServer 初始化http服务
func initHTTPServer(config *config.Config, handler http.Handler) {
	srv := &http.Server{
		Addr:         ":" + convert.ToString(config.Web.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
