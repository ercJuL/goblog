package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/rs/zerolog/log"
	"goblog/config"
	"goblog/model/response"
	"math/rand"
)

var store = base64Captcha.DefaultMemStore

// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	var driver base64Captcha.Driver
	switch rand.Intn(3) {
	case 0:
		driver = base64Captcha.NewDriverDigit(
			config.ServerConfig.Captcha.ImgHeight,
			config.ServerConfig.Captcha.ImgWidth,
			config.ServerConfig.Captcha.KeyLong,
			0.7,
			80)
	case 1:
		driver = base64Captcha.NewDriverString(
			config.ServerConfig.Captcha.ImgHeight,
			config.ServerConfig.Captcha.ImgWidth,
			0,
			base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine|base64Captcha.OptionShowSineLine,
			config.ServerConfig.Captcha.KeyLong,
			"1234567890qwertyuioplkjhgfdsazxcvbnm",
			nil, nil)
	case 2:
		driver = base64Captcha.NewDriverMath(
			config.ServerConfig.Captcha.ImgHeight,
			config.ServerConfig.Captcha.ImgWidth,
			0,
			base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine|base64Captcha.OptionShowSineLine,
			nil, nil)
	}

	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		log.Err(err).Stack().Msg("验证码获取失败")
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
