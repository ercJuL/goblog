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
	"goblog/utils"
)

func DBInit() {
	scm := config.ServerConfig.Ent
	db, err := sql.Open("pgx", scm.PostgreSqlDsn())
	if err != nil {
		log.Panic().Err(err).Msg("打开数据库连接错误")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	if scm.Debug {
		dao.Client = ent.NewClient(ent.Driver(drv), ent.Debug(), ent.Log(utils.EntLog))
	} else {
		dao.Client = ent.NewClient(ent.Driver(drv))
	}

	if err := dao.Client.Schema.Create(context.Background()); err != nil {
		log.Panic().Err(err).Msg("表结构更新失败")
	}

}

func DBDefer() {
	if err := dao.Client.Close(); err != nil {
		log.Err(err).Stack().Msg("关闭数据库连接错误")
	}
}
