package small_ant_dao

import (
	"small-ant-parent/small-ant-commn/database"
	model "small-ant-parent/small-ant-model"
)

type UserDao struct{}

func (c *UserDao) Info(where *model.User) (source model.User, notFound bool, err error) {
	notFound, err = database.First(&where, &source)
	return
}
