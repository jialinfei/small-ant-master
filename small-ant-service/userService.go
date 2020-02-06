package small_ant_service

import (
	"errors"
	"github.com/it234/goapp/pkg/logger"
	small_ant_dao "small-ant-parent/small-ant-dao"
	model "small-ant-parent/small-ant-model"
)

type UserService struct{}

var userDao = new(small_ant_dao.UserDao)

func (c *UserService) Info(u *model.User) (user model.User, err error) {
	user, notParam, err := userDao.Info(&model.User{Username: u.Username})
	if err != nil {
		if notParam {
			err = errors.New("没有找到此用户")
		}
		logger.Error(err)
		return user, err
	}
	return user, err
}
