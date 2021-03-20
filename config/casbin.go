package config

type Casbin struct {
	ModelPath string `mapstructure,json,yaml:"model_path"`
}
