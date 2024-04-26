package db

import (
	"fmt"
	"strings"
	"tiktok-lite/src/config"
	util "tiktok-lite/src/util/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"gorm.io/plugin/opentelemetry/tracing"
)

var Client *gorm.DB

func init() {
	var err error

	logger := util.GetGormLogger()

	cfg := gorm.Config{
		PrepareStmt: true,
		Logger:      logger,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4_general_ci&parseTime=True&loc=%s",
		config.EnvCfg.MySQL.User,
		config.EnvCfg.MySQL.Password,
		config.EnvCfg.MySQL.Host,
		config.EnvCfg.MySQL.Port,
		config.EnvCfg.MySQL.DB,
		config.EnvCfg.MySQL.TimeZone)

	if Client, err = gorm.Open(mysql.Open(dsn), &cfg); err != nil {
		// 连接数据库失败
		panic(err)
	}

	if config.EnvCfg.MySQL.Replica.Enable {
		var replicas []gorm.Dialector

		for _, addr := range strings.Split(config.EnvCfg.MySQL.Replica.Addr, ",") {
			pair := strings.Split(addr, ":")
			if len(pair) != 2 {
				continue
			}

			replicas = append(replicas, mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4_general_ci&parseTime=True&loc=%s",
				config.EnvCfg.MySQL.Replica.User,
				config.EnvCfg.MySQL.Replica.Password,
				pair[0],
				pair[1],
				config.EnvCfg.MySQL.Replica.DB,
				config.EnvCfg.MySQL.Replica.TimeZone)))
		}

		if err := Client.Use(dbresolver.Register(dbresolver.Config{
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		})); err != nil {
			panic(err)
		}
	}

	// 设置数据库连接池参数
	sqlDB, err := Client.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)                // 连接池中最大空闲连接数
	sqlDB.SetMaxOpenConns(100)               // 打开数据库的最大连接数
	sqlDB.SetConnMaxLifetime(24 * time.Hour) // 设置了连接最大存活时间
	sqlDB.SetConnMaxIdleTime(time.Hour)      // 设置空闲连接最大存活时间

	if err := Client.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

}
