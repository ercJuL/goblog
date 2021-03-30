package router

import (
	"github.com/gin-gonic/gin"
	v1 "goblog/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", v1.Register)
		UserRouter.POST("login", v1.Login)
		// UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
		// UserRouter.PUT("setUserInfo", v1.SetUserInfo)            // 设置用户信息
	}
}
