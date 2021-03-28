package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "goblog/docs"
	"goblog/middleware"
	"goblog/router"
)

func RouterInit() *gin.Engine {
	var gRouter = gin.Default()
	// e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	gRouter.Use(middleware.Cors())
	// gRouter.Use(authz.NewAuthorizer(e))
	gRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	PublicGroup := gRouter.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	return gRouter
}
