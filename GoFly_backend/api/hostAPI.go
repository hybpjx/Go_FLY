package api

import (
	"github.com/gin-gonic/gin"
	"gofly/response"
	"gofly/service"
	"gofly/service/dto"
)

type HostAPI struct {
	BaseAPI
	Service *service.HostService
}

func NewHostAPI() HostAPI {
	return HostAPI{
		Service: service.NewHostService(),
	}
}

func (h HostAPI) ShutDown(c *gin.Context) {
	var iShutDownHostDTO dto.ShutDownHostDTO
	if err := h.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iShutDownHostDTO,
		BindURI: false}); err != nil {
		return
	}
	err := h.Service.ShutDown(iShutDownHostDTO)
	if err != nil {
		h.Fail(response.Response{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}
	h.Success(response.Response{Msg: "shutdown ok"})
}
