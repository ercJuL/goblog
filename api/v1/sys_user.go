package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model/ent"
	"goblog/model/request"
	"goblog/model/response"
	"goblog/service"
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
