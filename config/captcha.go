package config

type Captcha struct {
	KeyLong   int `mapstructure,json,yaml:"key_long"`
	ImgWidth  int `mapstructure,json,yaml:"img_width"`
	ImgHeight int `mapstructure,json,yaml:"img_height"`
}
