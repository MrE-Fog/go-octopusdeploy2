package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type LibraryVariableSetUsageEntry struct {
	LibraryVariableSetID   string `json:"LibraryVariableSetId,omitempty"`
	LibraryVariableSetName string `json:"LibraryVariableSetName,omitempty"`

	Resource
}

// GetID returns the ID value of the LibraryVariableSetUsageEntry.
func (resource LibraryVariableSetUsageEntry) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this LibraryVariableSetUsageEntry.
func (resource LibraryVariableSetUsageEntry) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this LibraryVariableSetUsageEntry was changed.
func (resource LibraryVariableSetUsageEntry) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this LibraryVariableSetUsageEntry.
func (resource LibraryVariableSetUsageEntry) GetLinks() map[string]string {
	return resource.Links
}

func (resource LibraryVariableSetUsageEntry) SetID(id string) {
	resource.ID = id
}

func (resource LibraryVariableSetUsageEntry) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource LibraryVariableSetUsageEntry) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the LibraryVariableSetUsageEntry and returns an error if invalid.
func (resource LibraryVariableSetUsageEntry) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		return err
	}

	return nil
}

var _ ResourceInterface = &LibraryVariableSetUsageEntry{}
