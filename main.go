package main

import (
	"github.com/gin-gonic/gin"
	"imserver/config"
	"imserver/middleware"
	"imserver/middleware/jwt"
	"imserver/middleware/log"
	"imserver/middleware/redis"
	"imserver/router"
	"imserver/websocket"
)

func main() {
	config.Config = &config.Configuration{}
	app := gin.Default()
	app.Use(log.LoggerToFile())
	app.Use(middleware.Recover)
	app.GET("/ws", websocket.NewWebsocketHandler().Handler)

	app.Use(jwt.JwtAuth)
	router.InitRouter(app)
	//go func() {
	//	for {
	//		log.Logger.Infof("已注册通道：%s , 总数： %d", fmt.Sprintf("%v", websocket.WebSocketManage.UserMap), len(websocket.WebSocketManage.UserMap))
	//		time.Sleep(time.Duration(30) * time.Second)
	//	}
	//}()
	defer redis.RedisPool.Close()
	err := app.Run(":8090") // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		panic(err)
	}

}
