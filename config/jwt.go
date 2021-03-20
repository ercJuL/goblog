package config

type JWT struct {
	SigningKey  string `mapstructure,json,yaml:"signing_key"`
	ExpiresTime int64  `mapstructure,json,yaml:"expires_time"`
	BufferTime  int64  `mapstructure,json,yaml:"buffer_time"`
}
