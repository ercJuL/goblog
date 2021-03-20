package initialize

import (
	"github.com/rs/zerolog/log"
	"goblog/config"
	"goblog/dao"
	"goblog/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func DBInit() {
	scm := config.ServerConfig.Gorm
	postgreSqlConfig := postgres.Config{
		DSN: scm.PostgreSqlDsn(), // DSN data source name
	}
	db, err := gorm.Open(postgres.New(postgreSqlConfig), scm.GormConfig())
	if err != nil {
		log.Panic().Err(err).Msg("db 初始化失败")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(scm.MaxIdleConns)
	sqlDB.SetMaxOpenConns(scm.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	dao.DefaultDB = db

	// todo: 自动索引模型
	if err := db.AutoMigrate(&model.Post{}); err != nil {
		log.Panic().Err(err).Msg("表结构更新失败")
	}

}
