package config

import (
	"log"
	"tiktok-lite/src/util"

	"github.com/spf13/viper"
)

var EnvCfg envConfig

type envConfig struct {
	MySQL   MySQLConfig   `mapstructure:"mysql"`
	Logger  LoggerConfig  `mapstructure:"logger"`
	Pod     PodConfig     `mapstructure:"pod"`
	Tracing TracingConfig `mapstructure:"tracing"`
}

type MySQLConfig struct {
	DB       string `mapstructure:"db"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	TimeZone string `mapstructure:"timezone"`
	Replica  struct {
		Enable   bool   `mapstructure:"enable"`
		DB       string `mapstructure:"db"`
		Addr     string `mapstructure:"addr"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		TimeZone string `mapstructure:"timezone"`
	} `mapstructure:"replica"`
}

type LoggerConfig struct {
	Level           string `mapstructure:"level"`
	WithTranceState string `mapstructure:"with_trance_state"`
	Tied            string `mapstructure:"tied"`
}

type PodConfig struct {
	IP string `mapstructure:"ip"`
}

type TracingConfig struct {
	OtelState   string  `mapstructure:"otel_state"`
	OtelSampler float64 `mapstructure:"otel_sampler"`
	Endpoint    string  `mapstructure:"endpoint"`
}

func init() {
	// 读取配置文件
	viper.SetConfigName("config")           // 不需要文件扩展名
	viper.SetConfigType("yaml")             // 设置配置文件格式为 YAML
	viper.AddConfigPath(util.GetFilePath()) // 配置文件路径（这里是当前目录）

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	if err := viper.Unmarshal(&EnvCfg); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
}
