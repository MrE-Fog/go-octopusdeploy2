package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Machines defines a collection of machines with built-in support for paged
// results from the API.
type Machines struct {
	Items []Machine `json:"Items"`
	PagedResults
}

// Machine represents deployment targets (or machine).
type Machine struct {
	Name              string           `json:"Name,omitempty"`
	Thumbprint        string           `json:"Thumbprint,omitempty"`
	URI               string           `json:"Uri,omitempty" validate:"omitempty,uri"`
	IsDisabled        bool             `json:"IsDisabled"`
	EnvironmentIDs    []string         `json:"EnvironmentIds"`
	Roles             []string         `json:"Roles"`
	MachinePolicyID   string           `json:"MachinePolicyId,omitempty"`
	DeploymentMode    string           `json:"TenantedDeploymentParticipation,omitempty" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	TenantIDs         []string         `json:"TenantIds"`
	TenantTags        []string         `json:"TenantTags"`
	Status            string           `json:"Status,omitempty"`
	HealthStatus      string           `json:"HealthStatus,omitempty"`
	HasLatestCalamari bool             `json:"HasLatestCalamari"`
	StatusSummary     string           `json:"StatusSummary,omitempty"`
	IsInProcess       bool             `json:"IsInProcess"`
	Endpoint          *MachineEndpoint `json:"Endpoint" validate:"required"`
	OperatingSystem   string           `json:"OperatingSystem,omitempty"`
	ShellName         string           `json:"ShellName,omitempty"`
	ShellVersion      string           `json:"ShellVersion,omitempty"`
	Resource
}

// Validate returns a collection of validation errors against the machine's
// internal values.
func (machine *Machine) Validate() error {
	validate := validator.New()
	err := validate.Struct(machine)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}

		return err
	}

	return nil
}
