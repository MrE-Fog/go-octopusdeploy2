package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type tenantService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newTenantService(sling *sling.Sling, uriTemplate string) *tenantService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &tenantService{
		name:        serviceTenantService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s tenantService) getByProjectIDPath(id string) (string, error) {
	if isEmpty(id) {
		return emptyString, createInvalidParameterError(operationGetByProjectID, parameterID)
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterProjectID] = id

	return s.getURITemplate().Expand(values)
}

func (s tenantService) getClient() *sling.Sling {
	return s.sling
}

func (s tenantService) getName() string {
	return s.name
}

func (s tenantService) getPagedResponse(path string) ([]model.Tenant, error) {
	resources := []model.Tenant{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Tenants), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Tenants)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s tenantService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new Tenant.
func (s tenantService) Add(resource *model.Tenant) (*model.Tenant, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

// DeleteByID deletes the tenant that matches the input ID.
func (s tenantService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// GetAll returns all tenants. If none can be found or an error occurs, it
// returns an empty collection.
func (s tenantService) GetAll() ([]model.Tenant, error) {
	items := []model.Tenant{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the tenant that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s tenantService) GetByID(id string) (*model.Tenant, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Tenant), path)
	if err != nil {
		return nil, createResourceNotFoundError("tenant", "ID", id)
	}

	return resp.(*model.Tenant), nil
}

// GetByIDs returns the accounts that match the input IDs.
func (s tenantService) GetByIDs(ids []string) ([]model.Tenant, error) {
	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []model.Tenant{}, err
	}

	return s.getPagedResponse(path)
}

// GetByProjectID performs a lookup and returns all tenants with a matching
// project ID.
func (s tenantService) GetByProjectID(id string) ([]model.Tenant, error) {
	path, err := s.getByProjectIDPath(id)
	if err != nil {
		return []model.Tenant{}, nil
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns all tenants with a matching
// partial name.
func (s tenantService) GetByPartialName(name string) ([]model.Tenant, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.Tenant{}, nil
	}

	return s.getPagedResponse(path)
}

// Update modifies a tenant based on the one provided as input.
func (s tenantService) Update(resource model.Tenant) (*model.Tenant, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Tenant), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Tenant), nil
}

var _ ServiceInterface = &tenantService{}
