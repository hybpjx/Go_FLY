package service

import (
	"context"
	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/spf13/viper"
	"gofly/service/dto"
)

type HostService struct {
	BaseService
}

var hostService *HostService

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}

func (h HostService) ShutDown(iShutDownHostDTO dto.ShutDownHostDTO) error {
	var errResult error
	stHostIP := iShutDownHostDTO.HostIP
	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "ssh",
		User:       viper.GetString("ansible.user.name"),
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  stHostIP,
		ModuleName: "command",
		Args:       viper.GetString("ansible.shutdownHost.args"),
		ExtraVars: map[string]any{
			"ansible_password": viper.GetString("ansible.user.password"),
		},
	}

	_adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		StdoutCallback:    "oneline",
	}

	errResult = _adhoc.Run(context.TODO())

	return errResult
}
