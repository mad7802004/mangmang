package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mangmang/pkg/setting"
	"log"
)

var Orm *gorm.DB

func Setup() {
	var err error

	dbType := setting.DatabaseSetting.Type
	dbName := setting.DatabaseSetting.Name
	user := setting.DatabaseSetting.User
	password := setting.DatabaseSetting.Password
	host := setting.DatabaseSetting.Host

	Orm, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}
	Orm.LogMode(true)

	Orm.SingularTable(true)
	Orm.DB().SetMaxIdleConns(10)
	Orm.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	defer Orm.Close()
}

