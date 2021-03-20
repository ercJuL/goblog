package config

type Email struct {
	To       string `mapstructure,json,yaml:"to"`
	Port     int    `mapstructure,json,yaml:"port"`
	From     string `mapstructure,json,yaml:"from"`
	Host     string `mapstructure,json,yaml:"host"`
	IsSSL    bool   `mapstructure,json,yaml:"is_ssl"`
	Secret   string `mapstructure,json,yaml:"secret"`
	Nickname string `mapstructure,json,yaml:"nickname"`
}
