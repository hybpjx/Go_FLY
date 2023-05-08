package dao

import (
	"gofly/model"
	"gofly/service/dto"
)

type UserDao struct {
	BaseDao
}

var userDao *UserDao

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			NewBaseDao(),
		}
	}

	return userDao
}

// GetUserByName 获取用户的名称
func (u *UserDao) GetUserByName(stUsername string) (model.User, error) {
	var iUser model.User
	err := u.ORM.Model(&iUser).Where("name=?", stUsername).First(&iUser).Error
	return iUser, err
}

// GetUserByNameAndPassword 获取用户的名称和密码
func (u *UserDao) GetUserByNameAndPassword(stUsername, stPassword string) model.User {
	var iUser model.User
	u.ORM.Model(&iUser).Where("name=? and password=?", stUsername, stPassword).First(&iUser)
	return iUser
}

// CheckUserNameExist 检查用户是否存在
func (u *UserDao) CheckUserNameExist(stUserName string) bool {
	var total int64
	u.ORM.Model(&model.User{}).Where("name = ?", stUserName).Count(&total)

	return total > 0
}

// AddUser 添加用户
func (u *UserDao) AddUser(iUserDTO *dto.UserAddDTO) error {
	var iUser model.User
	// 构造一个用户实体
	iUserDTO.ConvertToModel(&iUser)
	err := u.ORM.Save(&iUser).Error

	if err == nil {
		iUserDTO.ID = iUser.ID
		iUserDTO.Password = ""
	}
	return err
}

// GetUserById 查询某一条用户
func (u *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User
	err := u.ORM.First(&iUser, id).Error
	return iUser, err
}

// GetUserList 查询所有的用户
func (u *UserDao) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	var iUserList []model.User
	var total int64

	err := u.ORM.Model(&model.User{}).
		Scopes(Paginate(iUserListDTO.Paginate)).
		Find(&iUserList).
		Offset(-1).
		Limit(-1).
		Count(&total).Error
	return iUserList, total, err
}

// UpdateUser 更新用户数据
func (u *UserDao) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	var iUser model.User

	u.ORM.First(&iUser, iUserUpdateDTO.ID)

	iUserUpdateDTO.ConvertToModel(&iUser)
	return u.ORM.Save(&iUser).Error
}

// DeleteUserByID 删除用户数据
func (u *UserDao) DeleteUserByID(id uint) error {

	return u.ORM.Delete(&model.User{}, id).Error
}
