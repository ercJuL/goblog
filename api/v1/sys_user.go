package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"goblog/config"
	"goblog/model"
	"goblog/model/ent"
	"goblog/model/request"
	"goblog/model/response"
	"goblog/service"
	"time"
)

// @Tags User
// @Summary 用户注册
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Register true "用户名, 昵称, 密码"
// @Success 200 {string} string "{"code":0,"data":{"id": 1},"msg":"创建用户成功"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var r request.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := ent.User{
		UserName:  r.Username,
		Password:  r.Password,
		NickName:  r.NickName,
		HeaderImg: r.HeaderImg,
	}
	user, err := service.Register(c, &u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithDetailed(response.Register{Id: user.ID}, "创建用户成功", c)
}

// @Tags User
// @Summary 用户注册
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Login true "用户名, 昵称, 密码"
// @Success 200 {string} string "{"code":0,"data":{"id": 1},"msg":"创建用户成功"}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var l request.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if !store.Verify(l.CaptchaId, l.Captcha, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}
	user, err := service.Login(c, l.Username, l.Password)
	if err != nil {
		response.FailWithMessage("用户名不存在或者密码错误", c)
	}
	tokenNext(c, user)
}

func tokenNext(c *gin.Context, user *ent.User) {
	nowUnix := time.Now().Unix()
	claims := model.JwtClaims{
		ID:        user.ID,
		Username:  user.UserName,
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: nowUnix + config.ServerConfig.JWT.ExpiresTime,
			Id:        "",
			IssuedAt:  nowUnix,
			Issuer:    "", // todo 想用域名作为提供者,但是多域名情况下会有问题
			NotBefore: nowUnix - 1000,
			Subject:   user.UserName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedString, err := token.SignedString(config.ServerConfig.System.SecretKey)
	if err != nil {
		log.Err(err).Stack().Msg("jwt签名错误")
	}
	response.OkWithDetailed(
		response.Login{
			Token:     signedString,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
		}, "登录成功", c,
	)
}
