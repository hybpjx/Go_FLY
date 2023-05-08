package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		// 允许所有站点均可访问
		//AllowAllOrigins: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return true
		//},

		// 允许所有站点均可访问 可将将 * 替换为指定的域名
		AllowOrigins: []string{"*"},
		//服务器支持的所有跨域请求的方法
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		//允许跨域设置可以返回其他子段，可以自定义字段
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Accept", "Token"},
		// 允许浏览器（客户端）可以解析的头部 （重要）
		ExposeHeaders: []string{"Content-Length"},
		//允许客户端传递校验信息比如 cookie (重要)
		AllowCredentials: true,
		//设置缓存时间
		MaxAge: 12 * time.Hour,
	})
}
