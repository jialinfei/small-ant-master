package small_ant_model

import (
	"small-ant-parent/small-ant-model/basemodel"
)

type User struct {
	Id       int64
	Username string `gorm:"size:50"`
	Password string `gorm:"size:50"`
}

func (User) TableName() string {
	return basemodel.TableName("sys_user")
}
