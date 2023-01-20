package dal

import (
	"github.com/chenshone/tiktok-lite/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	DB = DB.Debug()
	return err
}
