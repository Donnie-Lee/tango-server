package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"imserver/common/jwt"
	"imserver/common/response"
	"imserver/middleware/log"
	redisConn "imserver/middleware/redis"
	"imserver/models/loginModel"
	"imserver/service/account"
	"imserver/util"
	"strconv"
)

type AccountController struct {
}

const CheckCodeIntervalKey string = "LOGIN_CHECK_CODE_INTERVAL:"
const CheckCodeKey string = "LOGIN_CHECK_CODE:"
const CheckCodeInterval int = 60
const CheckCodeExpire int = 5 * 60

func GetCheckCode(context *gin.Context) {
	mobile := context.Param("mobile")
	redisConnection := redisConn.RedisPool.Get()
	checkCodeInterval, _ := redis.String(redisConnection.Do("get", CheckCodeIntervalKey+mobile))
	util.AssertExec(util.IsEmpty(checkCodeInterval), "获取验证码太频繁，请稍后再试")

	checkCode := util.RandInt(1000, 9999)
	log.Logger.Infof("获取验证码成功，手机号码:%s, 验证码: %d", mobile, checkCode)

	redis.String(redisConnection.Do("set", CheckCodeIntervalKey+mobile, checkCode, "EX", CheckCodeInterval))
	redis.String(redisConnection.Do("set", CheckCodeKey+mobile, checkCode, "EX", CheckCodeExpire))

	context.JSON(200, response.Success())
}

func LoginSms(ctx *gin.Context) {
	request := loginModel.LoginSmsRequest{}
	ctx.Bind(&request)
	redisConnection := redisConn.RedisPool.Get()
	checkCode, _ := redis.String(redisConnection.Do("get", CheckCodeKey+request.Mobile))
	util.AssertExec(util.IsNotEmpty(checkCode), "验证码已过期,请重新获取")
	util.AssertExec(checkCode == request.CheckCode, "验证码错误")

	redisConnection.Do("del", CheckCodeIntervalKey+request.Mobile)
	redisConnection.Do("del", CheckCodeKey+request.Mobile)

	account.LoginSms(request.Mobile, ctx)
}

func CurrenUser(ctx *gin.Context) {
	claims, _ := ctx.Get(`claims`)
	account.GetAccountInfo(claims.(jwt.UserClaims).LoginId, ctx)
}

func AccountInfo(ctx *gin.Context) {
	accountId, _ := strconv.Atoi(ctx.Param("accountId"))
	account.GetAccountInfo(accountId, ctx)
}

func Contacts(ctx *gin.Context) {
	account.Contacts(ctx)
}
