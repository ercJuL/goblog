package config

type System struct {
	Env           string `mapstructure,json,yaml:"env"`
	Addr          int    `mapstructure,json,yaml:"addr"`
	SecretKey     string `mapstructure,json,yaml:"secret_key"`
	DbType        string `mapstructure,json,yaml:"db_type"`
	OssType       string `mapstructure,json,yaml:"oss_type"`
	UseMultipoint bool   `mapstructure,json,yaml:"use_multipoint"`
}
