package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"imserver/config"
	"io"
	"os"
	"time"
)

var Logger *logrus.Logger

func LoggerToFile() gin.HandlerFunc {

	fileName := config.Log_FILE_PATH //写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		src, err = os.Create(fileName)
	}

	//实例化
	Logger = logrus.New()
	mw := io.MultiWriter(os.Stdout, src)
	//设置输出
	Logger.SetOutput(mw)

	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.Request.Host // 日志格式

		Logger.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
	}

}
