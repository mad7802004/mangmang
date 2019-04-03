package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mangmang/pkg/setting"
	"github.com/mangmang/pkg/utils"
	"log"
	"time"
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

	Orm.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	Orm.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	Orm.Callback().Create().Replace("gorm:before_create", defaultUuidCreateCallback)

}

func CloseDB() {
	defer Orm.Close()
}

func defaultUuidCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		Id := utils.GetUUID()
		if IdField, ok := scope.FieldByName("Id"); ok {
			if IdField.IsBlank {
				_ = IdField.Set(Id)
			}
		}
	}

}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdateTime"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}

		if BirthdayField, ok := scope.FieldByName("Birthday"); ok {
			if BirthdayField.IsBlank {
				_ = BirthdayField.Set(nowTime)
			}
		}

	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("UpdateTime", time.Now())
	}
}
