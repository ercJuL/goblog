package config

type Redis struct {
	DB       int    `mapstructure,json,yaml:"db"`
	Addr     string `mapstructure,json,yaml:"addr"`
	Password string `mapstructure,json,yaml:"password"`
}
