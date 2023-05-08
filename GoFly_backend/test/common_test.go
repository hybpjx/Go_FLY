package test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCommon(t *testing.T) {

	password, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	fmt.Println(string(password))

	// 可以解析出上文
	cost, _ := bcrypt.Cost([]byte("$2a$10$XgLBtSfJsrBd.liLOYWddOYWYWboBUAlKmivcSwq647C3vTNUOVMO"))
	fmt.Println(cost)

	err := bcrypt.CompareHashAndPassword(password, []byte("123"))
	if err != nil {
		fmt.Println("密码验证错误", err)
	}
	fmt.Println("密码验证成功>>>", nil)
}
