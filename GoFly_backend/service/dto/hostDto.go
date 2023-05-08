package dto

type ShutDownHostDTO struct {
	HostIP string `json:"host_ip" binding:"required" message:"必传参数"`
}
