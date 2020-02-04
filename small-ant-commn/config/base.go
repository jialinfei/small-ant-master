package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"small-ant-parent/small-ant-commn/database"
	"time"
)

func InitDB(config *Config) {
	var gdb *gorm.DB
	var err error
	if config.Gorm.DBType == "mysql" {
		config.Gorm.DSN = config.MySQL.DSN()
	} else if config.Gorm.DBType == "sqlite3" {
		config.Gorm.DSN = config.Sqlite3.DSN()
	}
	gdb, err = gorm.Open(config.Gorm.DBType, config.Gorm.DSN)
	if err != nil {
		panic(err)
	}
	gdb.SingularTable(true)
	if config.Gorm.Debug {
		gdb.LogMode(true)
		gdb.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}
	gdb.DB().SetMaxIdleConns(config.Gorm.MaxIdleConns)
	gdb.DB().SetMaxOpenConns(config.Gorm.MaxOpenConns)
	gdb.DB().SetConnMaxLifetime(time.Duration(config.Gorm.MaxLifetime) * time.Second)
	database.DB = gdb
}
