package variables

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type LibraryVariableSetUsageEntry struct {
	LibraryVariableSetID   string `json:"LibraryVariableSetId,omitempty"`
	LibraryVariableSetName string `json:"LibraryVariableSetName,omitempty"`

	resources.Resource
}

func NewLibraryVariableSetUsageEntry() *LibraryVariableSetUsageEntry {
	return &LibraryVariableSetUsageEntry{
		Resource: *resources.NewResource(),
	}
}
