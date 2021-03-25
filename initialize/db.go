package initialize

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
	"goblog/config"
	"goblog/dao"
	"goblog/model/ent"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Panic().Err(err).Msg("打开数据库连接错误")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func DBInit() {
	scm := config.ServerConfig.Ent
	dao.Client = Open(scm.PostgreSqlDsn())

	// todo: 自动索引模型
	if err := dao.Client.Schema.Create(context.Background()); err != nil {
		log.Panic().Err(err).Msg("表结构更新失败")
	}

}

func DBDefer() {
	if err := dao.Client.Close(); err != nil {
		log.Error().Err(err).Msg("关闭数据库连接错误")
	}
}
