package config

type Server struct {
	JWT        JWT        `mapstructure,json,yaml:"jwt"`
	ZeroLog    ZeroLog    `mapstructure,json,yaml:"zerolog"`
	Redis      Redis      `mapstructure,json,yaml:"redis"`
	Email      Email      `mapstructure,json,yaml:"email"`
	Casbin     Casbin     `mapstructure,json,yaml:"casbin"`
	System     System     `mapstructure,json,yaml:"system"`
	Captcha    Captcha    `mapstructure,json,yaml:"captcha"`
	Gorm       Gorm       `mapstructure,json,yaml:"gorm"`
	Local      Local      `mapstructure,json,yaml:"local"`
	Qiniu      Qiniu      `mapstructure,json,yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure,json,yaml:"aliyun_oss"`
	TencentCOS TencentCOS `mapstructure,json,yaml:"tencent_cos"`
	Excel      Excel      `mapstructure,json,yaml:"excel"`
}

var ServerConfig Server
