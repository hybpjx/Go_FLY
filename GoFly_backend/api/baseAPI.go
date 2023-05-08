package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gofly/global"
	"gofly/response"
	"gofly/utils"
	"reflect"
)

type BaseAPI struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

type BaseAPIer interface {
	AddError(errNew error)
	GetError() error
	BuildRequest(option BuildRequestOption) *BaseAPI
	ParseValidateErrors(errs error, target any) error
	Fail(resp response.Response)
	Success(resp response.Response)
	ServerFail(resp response.Response)
}

// BuildRequestOption 构建请求选项option
type BuildRequestOption struct {
	Ctx     *gin.Context
	DTO     any
	BindURI bool
	BindAny bool
}

// NewBaseAPI 初始化BaseAPI
func NewBaseAPI() BaseAPI {
	return BaseAPI{
		Logger: global.Logger,
	}
}

// AddError 添加错误 因为需要改变层面属性 所以必须要制定指针
func (b *BaseAPI) AddError(errNew error) {
	b.Errors = utils.AppendError(b.Errors, errNew)

}

// GetError 获取错误 他的作用就是返回经过调用的BaseAPI中的error
func (b *BaseAPI) GetError() error {
	return b.Errors
}

func (b *BaseAPI) BuildRequest(option BuildRequestOption) *BaseAPI {
	var errorResult error
	// 绑定请求上下文
	b.Ctx = option.Ctx
	// 绑定请求数据
	if option.DTO != nil {
		// 是否绑定URI中的参数
		if option.BindAny || option.BindURI {
			errorResult = utils.AppendError(errorResult, b.Ctx.ShouldBindUri(option.DTO))
		}
		if option.BindAny || !option.BindURI {
			errorResult = utils.AppendError(errorResult, b.Ctx.ShouldBind(option.DTO))
		}

		if errorResult != nil {
			errorResult = b.ParseValidateErrors(errorResult, option.DTO)
			b.AddError(errorResult)

			b.Fail(response.Response{
				Msg: b.GetError().Error()},
			)

		}

	}
	return b
}

// ParseValidateErrors 解析错误类型 而传入的类型 必须是ValidationErrors类型的
// 还需要传一个对象 最终目标就是获取tag里的参数 从而完成错误的解析
func (b *BaseAPI) ParseValidateErrors(errs error, target any) error {
	var errorResult error

	// 做一个错误的类型断言
	errValidations, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}

	// 通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidations {
		field, _ := fields.FieldByName(fieldErr.Field())
		//  {Name  string json:"name" binding:"required,first_test" message:"用户名填写错误" required_err:"用户名不能为空" 0 [0] false}
		//fmt.Println("field>>>>>>>>", field)

		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		//  required_err
		//fmt.Println("errMessageTag>>>>>>>>>>", errMessageTag)
		errMessage := field.Tag.Get(errMessageTag)

		// 用户名不能为空
		//fmt.Println("errMessage>>>>>>>>>>", errMessage)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s:%s Error", fieldErr.Error(), fieldErr.Tag())
		}

		errorResult = utils.AppendError(errorResult, errors.New(errMessage))

	}
	return errorResult
}

// Fail 包装好的失败的返回
func (b *BaseAPI) Fail(resp response.Response) {
	response.Fail(b.Ctx, resp)
}

// Success 包装好的成功的返回
func (b *BaseAPI) Success(resp response.Response) {
	response.Success(b.Ctx, resp)
}

// ServerFail 包装500Server的返回
func (b *BaseAPI) ServerFail(resp response.Response) {
	response.ServerFail(b.Ctx, resp)
}
