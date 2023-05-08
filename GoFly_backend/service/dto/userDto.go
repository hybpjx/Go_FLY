package dto

import (
	"gofly/model"
)

// UserLoginDTO 该文件是针对于用户登录的dto
type UserLoginDTO struct {
	// 还可以针对bing校验 来抛错，message 是默认所有的错，required_err 是针对与 required的错误
	Name     string `json:"name" binding:"required" message:"用户名填写错误"`
	Password string `json:"password" binding:"required" message:"用户名密码不能为空"`
}

// UserAddDTO ===== 添加用户相关DTO =====
type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	// 如果为空就不接受了
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"密码不能为空"`
}

func (u UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Name = u.Name
	iUser.RealName = u.RealName
	iUser.Mobile = u.Mobile
	iUser.Avatar = u.Avatar
	iUser.Email = u.Email

	// 通过方法直接调用
	//hashPassword, _ := utils.Encrypt(u.Password)
	//iUser.Password = hashPassword

	iUser.Password = u.Password
}

// UserUpdateDTO ===== 更新用户 =====
type UserUpdateDTO struct {
	ID       uint   `json:"id" form:"id" uri:"id"`
	Name     string `json:"name" form:"name"`
	RealName string `json:"real_name" form:"real_name"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

func (u UserUpdateDTO) ConvertToModel(iUser *model.User) {
	iUser.ID = u.ID
	iUser.Name = u.Name
	iUser.RealName = u.RealName
	iUser.Mobile = u.Mobile
	iUser.Email = u.Email
}

// UserListDTO ===== 用户列表相关的DTO =====
type UserListDTO struct {
	Paginate
}
