package proxies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type ProxyService struct {
	services.CanDeleteService
}

func NewProxyService(sling *sling.Sling, uriTemplate string) *ProxyService {
	return &ProxyService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceProxyService, sling, uriTemplate),
		},
	}
}
