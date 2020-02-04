package small_ant_service

import model "small-ant-parent/small-ant-model"

type UserService interface {
	Info() *model.User
}

func Info(u *model.User) model.User {
	model := model.User{}
	return model
}
