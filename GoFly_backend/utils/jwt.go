package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

// 把签发的秘钥 抛出来
var stSignKey = []byte(viper.GetString("jwt.SignKey"))

// JwtCustomClaims 注册声明是JWT声明集的结构化版本，仅限于注册声明名称
type JwtCustomClaims struct {
	ID               uint
	Name             string
	RegisteredClaims jwt.RegisteredClaims
}

func (j JwtCustomClaims) Valid() error {
	return nil
}

// GenerateToken 生成Token
func GenerateToken(id uint, name string) (string, error) {
	// 初始化
	iJwtCustomClaims := JwtCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间 在当前基础上 添加一个小时后 过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.TokenExpire") * time.Minute)),
			// 颁发时间 也就是生成时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//主题
			Subject: "Token",
		},
	}
	/*
		使用HS256 的签名加密方法
		使用指定的签名方法和声明创建一个新的[Token]
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	return token.SignedString(stSignKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	// 声明一个空的数据声明
	iJwtCustomClaims := JwtCustomClaims{}
	//ParseWithClaims是NewParser().ParseWithClaims()的快捷方式
	//第一个值是token ，
	//第二个值是我们之后需要把解析的数据放入的地方，
	//第三个值是Keyfunc将被Parse方法用作回调函数，以提供用于验证的键。函数接收已解析但未验证的令牌。
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})

	// 判断 是否为空 或者是否无效只要两边有一处是错误 就返回无效token
	if err == nil && !token.Valid {
		err = errors.New("invalid Token")
	}
	return iJwtCustomClaims, err
}

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	fmt.Println(err)
	if err != nil {
		return false
	}
	return true
}
