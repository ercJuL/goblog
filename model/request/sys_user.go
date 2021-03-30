package request

type Register struct {
	Username  string `json:"userName"`
	Password  string `json:"passWord"`
	NickName  string `json:"nickName"`
	HeaderImg string `json:"headerImg"`
}

type Login struct {
	Username  string `json:"userName"`
	Password  string `json:"passWord"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
