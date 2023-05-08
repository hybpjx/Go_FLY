package service

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gofly/dao"
	"gofly/global"
	"gofly/global/constants"
	"gofly/model"
	"gofly/service/dto"
	"gofly/utils"
	"strconv"
	"strings"
	"time"
)

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

var userService *UserService

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

// GenerateAndCacheLoginUserToken 向redis写入 token信息 每三十秒或者三十分钟就过期
func GenerateAndCacheLoginUserToken(userID uint, stUserName string) (string, error) {
	token, err := utils.GenerateToken(userID, stUserName)
	if err == nil {
		err = global.RedisClient.Set(strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(userID)), -1), token, viper.GetDuration("jwt.TokenExpire")*time.Minute)
	}
	return token, err
}

// Login 登录
func (u UserService) Login(iUserDTo dto.UserLoginDTO) (model.User, string, error) {
	// 定义一个错误
	var errResult error

	// 定义一个token字符串
	var token string

	iUser, err := u.Dao.GetUserByName(iUserDTo.Name)

	// 用户名或者密码不正确
	if (err != nil) || !(utils.CompareHashAndPassword(iUser.Password, iUserDTo.Password)) {
		errResult = errors.New("invalid UserName or Password")
	} else {
		// 登录成功  即验证成功 ——> 生成Token
		token, err = GenerateAndCacheLoginUserToken(iUser.ID, iUser.Name)
		if err != nil {
			errResult = errors.New(fmt.Sprintf("generate Token errors:%s", err.Error()))
		}
	}

	return iUser, token, errResult
}

func (u UserService) AddUser(iUserDTO *dto.UserAddDTO) error {
	// 如果返回的是true  说明已存在
	if u.Dao.CheckUserNameExist(iUserDTO.Name) {
		return errors.New("user Name Exist")
	}
	return u.Dao.AddUser(iUserDTO)
}

func (u *UserService) GetUserByID(iCommonIDDTO *dto.CommonIDDTO) (model.User, error) {
	return u.Dao.GetUserById(iCommonIDDTO.ID)
}

// GetUserList 获取所有用户
func (u *UserService) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return u.Dao.GetUserList(iUserListDTO)
}

// UpdateUser 更新某一条用户数据
// UpdateUser 更新用户数据
func (u *UserService) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	if iUserUpdateDTO.ID == 0 {
		return errors.New("invalid User ID")
	}

	return u.Dao.UpdateUser(iUserUpdateDTO)
}

// DeleteUserByID 删除某一条用户数据
// DeleteUserByID 删除用户数据
func (u *UserService) DeleteUserByID(iCommonIDDTO *dto.CommonIDDTO) error {
	if iCommonIDDTO.ID == 0 {
		return errors.New("invalid User ID")
	}

	return u.Dao.DeleteUserByID(iCommonIDDTO.ID)
}
