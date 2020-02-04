package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/it234/goapp/pkg/logger"
	"small-ant-parent/small-ant-commn/database"
	"small-ant-parent/small-ant-commn/response.common"
	model "small-ant-parent/small-ant-model"
)

func Info(c *gin.Context) {
	username := response_common.GetQueryToStr(c, "username")
	password := response_common.GetQueryToStr(c, "password")
	if username == "" || password == "" {
		response_common.ResFail(c, "用户名或密码不能为空")
		return
	}
	where := model.User{Username: username}
	user := model.User{}
	notFound, err := database.First(&where, &user)
	if err != nil {
		if notFound {
			response_common.ResFail(c, "没用此用户")
			return
		}
		response_common.ResErrSrv(c, err)
		logger.Error(err)
		return
	}
	response_common.ResSuccess(c, &user)
}

func Login(c *gin.Context) {
	requestData, err := c.GetRawData()
	if err != nil {
		response_common.ResErrSrv(c, err)
		return
	}
	var requestMap map[string]string
	err = json.Unmarshal(requestData, &requestMap)
	if err != nil {
		response_common.ResErrSrv(c, err)
		return
	}
	username := requestMap["username"]
	password := requestMap["password"]
	if username == "" || password == "" {
		response_common.ResFail(c, "用户名或密码不能为空")
		return
	}
	where := model.User{Username: "张山"}
	user := model.User{}
	notFound, err := database.First(&where, &user)
	if err != nil {
		if notFound {
			response_common.ResFail(c, "没用此用户")
			return
		}
		response_common.ResErrSrv(c, err)
		logger.Error(err)
		return
	}
	response_common.ResSuccess(c, &user)
}
