package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gofly/response"
	"gofly/service"
	"gofly/service/dto"
	"net/http"
)

const (
	ErrorAddUserCode        = 10011
	ErrorGetUserByIDCode    = 10012
	ErrorGetUserListCode    = 10013
	ErrorUpdateUserCode     = 10014
	ErrorDeleteUserByIDCode = 10015
	ErrorLoginCOde          = 10016
)

type UserAPI struct {
	BaseAPI
	Service *service.UserService
}

func NewUserAPI() UserAPI {
	return UserAPI{
		BaseAPI: NewBaseAPI(),
		Service: service.NewUserService(),
	}
}

// Login 用户登录方法
// @Tag 用户管理
// @Summary 用户登录
// @Description 用户登录详情表述
// @Accept  json
// @Produce  json
// @Param  username formData string true "用户名"
// @Param  password formData string true "密码"
// @Success 200 {string} string	"登入成功"
// @Failure 401 {string} string "登入失败"
// @Router /api/v1/public/user/login [post]
func (u UserAPI) Login(ctx *gin.Context) {

	// 声明一个 userDTO类型 用来校验用户名和密码
	var iUserLoginDTO dto.UserLoginDTO
	//errs := ctx.ShouldBind(&iUserLoginDTO)
	//fmt.Println("errs>>>>", errs)
	//NewUserAPI()
	//// 如果错误 就返回错误信息
	//if errs != nil {
	//
	//	response.Fail(ctx, response.Response{
	//		Msg: u.ParseValidateErrors(errs.(validator.ValidationErrors), &iUserLoginDTO).Error(),
	//	})
	//	return
	//}

	// 构建一个请求绑定 绑定上下文和 DTO校验信息 然后 GetError获取错误信息
	if err := u.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	iUser, token, err := u.Service.Login(iUserLoginDTO)
	if err != nil {
		u.Fail(response.Response{
			Code:   ErrorLoginCOde,
			Msg:    err.Error(),
			Status: http.StatusUnauthorized,
		})
		return
	}
	//token, err = service.GenerateAndCacheLoginUserToken(iUser.ID, token)

	u.Success(response.Response{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
		Msg: "登录成功",
	})

}

func (u UserAPI) AddUser(c *gin.Context) {
	// 声明一个结构体 由于绑定参数和校验
	var iUserAddDTO dto.UserAddDTO
	err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO, BindURI: false}).GetError()
	if err != nil {
		fmt.Println(err)
		return
	}

	////上传用户头绪
	//file, err := c.FormFile("file")
	//if err != nil {
	//	fmt.Println("文件读取失败")
	//	panic(err)
	//}
	//stFilePath := fmt.Sprintf("./upload/%s", file.Filename)
	//if err = c.SaveUploadedFile(file, stFilePath); err != nil {
	//	fmt.Println("保存失败")
	//	panic(err)
	//}
	//iUserAddDTO.Avatar = stFilePath

	// 其次 把针对数据库操作的封装到services中
	err = u.Service.AddUser(&iUserAddDTO)

	// 返回错误
	if err != nil {
		u.ServerFail(response.Response{
			Code: ErrorAddUserCode,
			Msg:  err.Error()},
		)
		return
	}

	// 返回成功
	u.Success(response.Response{Data: iUserAddDTO})
}

// GetUserByID 通过ID 获取user信息
func (u UserAPI) GetUserByID(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindURI: true}).GetError(); err != nil {
		return
	}
	iUser, err := u.Service.GetUserByID(&iCommonIDDTO)
	//iUser.Password = "******"
	if err != nil {
		u.ServerFail(response.Response{
			Code: ErrorGetUserByIDCode,
			Msg:  err.Error(),
		})
		return
	}

	u.Success(response.Response{Data: iUser})
}

func (u UserAPI) GetUserList(c *gin.Context) {
	var iUserListDTO dto.UserListDTO
	if err := u.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iUserListDTO,
		BindURI: false,
	}).GetError(); err != nil {

		return
	}

	giUserList, nTotal, err := u.Service.GetUserList(&iUserListDTO)
	if err != nil {
		u.ServerFail(response.Response{
			Code: ErrorGetUserListCode,
			Msg:  err.Error(),
		})
		return
	}

	u.Success(response.Response{
		Data:  giUserList,
		Total: nTotal,
	})
}

// UpdateUser 更新某条用户
func (u UserAPI) UpdateUser(c *gin.Context) {
	var iUserUpdateDTO dto.UserUpdateDTO

	//paramsID := c.Param("id")
	//id, _ := strconv.Atoi(paramsID)
	//uid := uint(id)
	//iUserUpdateDTO.ID = uid

	if err := u.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iUserUpdateDTO,
		BindAny: true,
	}).GetError(); err != nil {

		return
	}

	if err := u.Service.UpdateUser(&iUserUpdateDTO); err != nil {
		u.ServerFail(response.Response{
			Code: ErrorUpdateUserCode,
			Msg:  err.Error(),
		})
		return
	}

	u.Success(response.Response{Msg: "修改成功", Data: iUserUpdateDTO})
}

func (u UserAPI) DeleteUserByID(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindURI: true}).GetError(); err != nil {
		return
	}
	err := u.Service.DeleteUserByID(&iCommonIDDTO)
	//iUser.Password = "******"
	if err != nil {
		u.ServerFail(response.Response{
			Code: ErrorDeleteUserByIDCode,
			Msg:  err.Error(),
		})
		return
	}

	u.Success(response.Response{Msg: "删除成功"})
}
