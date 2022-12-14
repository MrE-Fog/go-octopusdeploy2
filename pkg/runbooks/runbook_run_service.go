package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type RunbookRunService struct {
	services.CanDeleteService
}

func NewRunbookRunService(sling *sling.Sling, uriTemplate string) *RunbookRunService {
	return &RunbookRunService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceRunbookRunService, sling, uriTemplate),
		},
	}
}
