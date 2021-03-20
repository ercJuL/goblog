package main

import (
	"github.com/rs/zerolog/log"
	"goblog/initialize"
)

func main() {
	initialize.ViperInit()
	initialize.ZeroLogInit()
	initialize.DBInit()
	router := initialize.RouterInit()
	router.Run("127.0.0.1:13831")
	log.Info().Msg("test")
	// dao.InitDB()
	// r := gin.New()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
