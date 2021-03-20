package config

type Local struct {
	Path string `mapstructure,json,yaml:"path" `
}

type Qiniu struct {
	Zone          string `mapstructure,json,yaml:"zone"`
	Bucket        string `mapstructure,json,yaml:"bucket"`
	ImgPath       string `mapstructure,json,yaml:"img_path"`
	UseHTTPS      bool   `mapstructure,json,yaml:"use_https"`
	AccessKey     string `mapstructure,json,yaml:"access_key"`
	SecretKey     string `mapstructure,json,yaml:"secret_key"`
	UseCdnDomains bool   `mapstructure,json,yaml:"use_cdn_domains"`
}

type AliyunOSS struct {
	Endpoint        string `mapstructure,json,yaml:"endpoint"`
	AccessKeyId     string `mapstructure,json,yaml:"access_key_id"`
	AccessKeySecret string `mapstructure,json,yaml:"access_key_secret"`
	BucketName      string `mapstructure,json,yaml:"bucket_name"`
	BucketUrl       string `mapstructure,json,yaml:"bucket_url"`
}
type TencentCOS struct {
	Bucket     string `mapstructure,json,yaml:"bucket"`
	Region     string `mapstructure,json,yaml:"region"`
	SecretID   string `mapstructure,json,yaml:"secret_id"`
	SecretKey  string `mapstructure,json,yaml:"secret_key"`
	BaseURL    string `mapstructure,json,yaml:"base_url"`
	PathPrefix string `mapstructure,json,yaml:"path_prefix"`
}
