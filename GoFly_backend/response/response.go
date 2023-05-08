package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Response struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  any    `json:"total,omitempty"`
}

// IsEmpty 判断结构体是否为空
func (r Response) IsEmpty() bool {
	return reflect.DeepEqual(r, Response{})
}

// 构建状态码 ，如果 传入的ResponseJson没有Status 就使用默认的状态码
func buildStatus(resp Response, defaultStatus int) int {
	if resp.Status == 0 {
		return defaultStatus
	}
	return resp.Status
}

func HttpResponse(ctx *gin.Context, status int, resp Response) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(status, resp)
}

func Success(ctx *gin.Context, resp Response) {
	HttpResponse(ctx, buildStatus(resp, http.StatusOK), resp)
}

func Fail(ctx *gin.Context, resp Response) {
	HttpResponse(ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

func ServerFail(ctx *gin.Context, resp Response) {
	HttpResponse(ctx, buildStatus(resp, http.StatusInternalServerError), resp)

}
