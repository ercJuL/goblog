package initialize

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"goblog/middleware"
)

func RouterInit() *gin.Engine {
	var router = gin.Default()
	e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	router.Use(middleware.Cors())
	router.Use(authz.NewAuthorizer(e))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
