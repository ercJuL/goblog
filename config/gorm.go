package config

import (
	"fmt"
	"goblog/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Gorm struct {
	Host         string `mapstructure,json,yaml:"host"`
	Port         int    `mapstructure,json,yaml:"port"`
	Params       string `mapstructure,json,yaml:"params"`
	Dbname       string `mapstructure,json,yaml:"db_name"`
	Username     string `mapstructure,json,yaml:"username"`
	Password     string `mapstructure,json,yaml:"password"`
	MaxIdleConns int    `mapstructure,json,yaml:"max_idle_conns"`
	MaxOpenConns int    `mapstructure,json,yaml:"max_open_conns"`
	LogMode      bool   `mapstructure,json,yaml:"log_mode"`
	LogLevel     string `mapstructure,json,yaml:"log_level"`
}

func (m *Gorm) MysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.Dbname, m.Params)
}

func (m *Gorm) PostgreSqlDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Shanghai", m.Host, m.Username, m.Password, m.Dbname, m.Port)
}

func (m *Gorm) GormConfig() *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch m.LogLevel { // 初始化配置文件的Level
	case "silent", "Silent":
		config.Logger = utils.Gormlogger.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = utils.Gormlogger.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = utils.Gormlogger.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = utils.Gormlogger.LogMode(logger.Info)
	}

	return config
}
