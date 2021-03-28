package config

import (
	"fmt"
)

type Ent struct {
	Host         string `mapstructure,json,yaml:"host"`
	Port         int    `mapstructure,json,yaml:"port"`
	Params       string `mapstructure,json,yaml:"params"`
	Dbname       string `mapstructure,json,yaml:"db_name"`
	Username     string `mapstructure,json,yaml:"username"`
	Password     string `mapstructure,json,yaml:"password"`
	MaxIdleConns int    `mapstructure,json,yaml:"max_idle_conns"`
	MaxOpenConns int    `mapstructure,json,yaml:"max_open_conns"`
	Debug        bool   `mapstructure,json,yaml:"debug"`
}

func (e *Ent) MysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", e.Username, e.Password, e.Host, e.Port, e.Dbname, e.Params)
}

func (e *Ent) PostgreSqlDsn() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", e.Username, e.Password, e.Host, e.Port, e.Dbname)
}
