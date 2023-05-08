package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gofly/global"
	"gofly/global/constants"
	"gofly/model"
	"gofly/response"
	"gofly/service"
	"gofly/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ERR_CODE_INVALID_TOKEN     = 10401 // token 无效 格式错误
	ERR_CODE_TOKEN_PARSE       = 10402 //解析token失败
	ERR_CODE_TOKEN_NOT_MATCHED = 10403 //TOKEN 访问者登录时Token不一致
	ERR_CODE_TOKEN_EXPIRED     = 10404 //TOKEN 已过期
	ERR_CODE_TOKEN_RENEW       = 10405 //token续期失败
	TOKEN_NAME                 = "Authorization"
	TOKEN_PREFIX               = "Bearer: "
	RENEW_TOKEN_DURATION       = 10 * 60 * time.Second
)

func tokenERR(c *gin.Context, code int) {
	response.Fail(c, response.Response{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "Invalid Token",
	})
}

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 取传过来的token
		token := c.GetHeader(TOKEN_NAME)
		// 判断token格式
		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			tokenERR(c, ERR_CODE_INVALID_TOKEN)
			fmt.Println("判断token格式错误")
			return
		}

		// token解析失败无法返回
		token = token[len(TOKEN_PREFIX):]
		iJwtCustomClaims, eRR := utils.ParseToken(token)
		nUserID := iJwtCustomClaims.ID
		if eRR != nil || nUserID == 0 {
			tokenERR(c, ERR_CODE_TOKEN_PARSE)
			fmt.Println("token解析失败无法返回")
			return
		}
		// Token 与访问者对应的token 不一致 直接返回
		stUserID := strconv.Itoa(int(nUserID))
		stRedisUserIDKey := strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", stUserID, -1)

		stRedisToken, eRR := global.RedisClient.Get(stRedisUserIDKey)
		if eRR != nil || stRedisToken != token {
			tokenERR(c, ERR_CODE_TOKEN_NOT_MATCHED)
			fmt.Println("Token 与访问者对应的token 不一致 直接返回")
			return
		}
		// Token 已经过期
		nTokenDuration, eRR := global.RedisClient.GetExpireDuration(stRedisUserIDKey)
		if eRR != nil || nTokenDuration <= 0 {
			tokenERR(c, ERR_CODE_TOKEN_EXPIRED)
			fmt.Println("Token 已经过期")
			return
		}

		// token的续期 代表客户端一只存活
		if nTokenDuration.Seconds() < RENEW_TOKEN_DURATION.Seconds() {
			// 1. 重新生成token
			stNewToken, eRR := service.GenerateAndCacheLoginUserToken(nUserID, iJwtCustomClaims.Name)
			if eRR != nil {
				tokenERR(c, ERR_CODE_TOKEN_RENEW)
				fmt.Println("token的续期 代表客户端一只存活")
				return
			}
			// 2. 把新的token 扔到token中去
			// 3. 把token 返回给客户端
			c.Header("token", stNewToken)
		}
		//iUser, eRR := dao.NewUserDao().GetUserById(nUserID)
		//if eRR != nil {
		//	tokenERR(c)
		//	return
		//}
		//c.Set(constants.LOGIN_USER, iUser)
		c.Set(constants.LOGIN_USER, model.LoginUser{
			ID:   nUserID,
			Name: iJwtCustomClaims.Name,
		})
		c.Next()

	}
}
