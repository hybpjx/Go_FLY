package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gofly/docs"
	"gofly/global"
	"gofly/middleware"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// IFnRegisterRoute 函数类型 通过这个方法去注册路由 第一个不需要做鉴权，第二个需要做鉴权认证 传token
type IFnRegisterRoute func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

// 按照模块来配置路由
var (
	gfnRoutes []IFnRegisterRoute
)

// RegisterRoute 注册路由回调函数
func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	// 如果不为空 就把这个函数往路由组里塞
	gfnRoutes = append(gfnRoutes, fn)
}

// InitRouter 初始化系统路由组
func InitRouter() {
	// ==============创建一个可被取消的ctx 去监听两个事件，分别是 ctrl c 和 退出应用 只要触发就进入结束状态 然后跳到 <-Done()的位置==============
	ctx, cancelCtx := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)

	defer cancelCtx()

	// ==============初始化Gin 框架==============
	r := gin.Default()

	r.Use(middleware.Cors())

	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1") // 鉴权的太多了 直接就用这个了
	rgAuth.Use(middleware.Auth())

	// ==============初始化基础路由模块==============
	// InitBasePlatformRoutes -> InitUserRoutes > RegisterRoute
	// 执行完这个方法 gfnRoutes 才会添加完数据 如果不执行这个方法 这就是个空数组
	initBasePlatformRoutes()

	// ==============注册初始化自定义验证器==============
	registerCustomValidator()

	//遍历所有的基础路由模块
	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	// ==============集成swagger==============
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ==============读取配置==============
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}

	// 初始化服务
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	go func() {
		global.Logger.Infof("Start Listen:%s", stPort)
		// 启动 server服务
		err := server.ListenAndServe()
		if (err != nil) && (err != http.ErrServerClosed) {
			global.Logger.Errorf("start server Error:%s", err.Error())
		}
		//fmt.Println(fmt.Sprintf("start server Listen:%s\n", stPort))
	}()

	// 读出来是个空的结构体 {} 然后就开始放行
	<-ctx.Done()
	// 有一个五秒钟的延迟，给服务5s 的时间去做关闭 如果出现错误 就记录日志 没有就继续往下就继续执行。
	ctx, cancelShutDown := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelShutDown()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Errorf("Stop server Error:%s", err.Error())
		return
	}
	// 走到这一步 main() 方法中的Start方法 真正意义的结束 然后开始执行clean方法
	global.Logger.Infoln("Stop Server Success")
}

// initBasePlatformRoutes 初始化基础路由模块组
// 各个业务的模块 比如用户模块
func initBasePlatformRoutes() {
	InitUserRoutes()
	InitHostRoutes()
}

// registerCustomValidator 用于注册 自定义验证器 比如 email字符串 身份证字符串的校验
func registerCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("first_test", func(fl validator.FieldLevel) bool {
			// 进行类型断言 断言成字符串 如果成功 才进行下一步操作。
			if value, ok := fl.Field().Interface().(string); ok {
				// 如果这个值不为空 且第一个值为t 就返回False
				if value != "" && strings.Index(value, "t") == 0 {
					return true
				}
			}
			return false
		})
	}

}
