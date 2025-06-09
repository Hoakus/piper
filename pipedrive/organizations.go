package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

type OrganizationPiper piper

type OrganizationResponse struct {
	Success bool         `json:"success"`
	Data    Organization `json:"data"`
}

type OrganizationsResponse struct {
	Success        bool            `json:"success"`
	Data           []Organization  `json:"data"`
	AdditionalData *AdditionalData `json:"additional_data,omitempty"`
}

// https://developers.pipedrive.com/docs/api/v1/Organizations#getOrganizations
func (piper *OrganizationPiper) GetAll(ctx context.Context, params OrganizationsGetOptions) (*OrganizationsResponse, *http.Response, error) {
	endpoint := "api/v2/organizations"

	request, err := piper.client.NewRequest("GET", endpoint, &params, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationsResponse

	response, err := piper.client.Do(ctx, request, &record)
	if err != nil {
		return nil, response, err
	}

	return record, response, nil
}

// https://developers.pipedrive.com/docs/api/v1/Organizations#getOrganization
func (piper *OrganizationPiper) Get(ctx context.Context, orgID int, params OrganizationGetOptions) (*OrganizationResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("api/v2/organizations/%d", orgID)

	request, err := piper.client.NewRequest("GET", endpoint, params, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationResponse

	response, err := piper.client.Do(ctx, request, &record)
	if err != nil {
		return nil, response, err
	}

	return record, response, nil
}

// https://developers.pipedrive.com/docs/api/v1/Organizations#addOrganization
func (piper *OrganizationPiper) Add(ctx context.Context, body OrganizationAddOptions) (*OrganizationResponse, *http.Response, error) {
	endpoint := "api/v2/organizations"

	request, err := piper.client.NewRequest("POST", endpoint, nil, body)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationResponse

	response, err := piper.client.Do(ctx, request, &record)
	if err != nil {
		return nil, response, err
	}

	return record, response, nil
}

// https://developers.pipedrive.com/docs/api/v1/Organizations#updateOrganization
func (piper *OrganizationPiper) Update(ctx context.Context, orgID int, body OrganizationUpdateOptions) (*OrganizationResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("api/v2/organizations/%d", orgID)

	request, err := piper.client.NewRequest("PATCH", endpoint, nil, body)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationResponse

	response, err := piper.client.Do(ctx, request, &record)
	if err != nil {
		return nil, response, err
	}

	return record, response, nil
}

// https://developers.pipedrive.com/docs/api/v1/Organizations#deleteOrganization
func (piper *OrganizationPiper) Delete(ctx context.Context, orgID int) (*OrganizationResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("api/v2/organizations/%d", orgID)

	request, err := piper.client.NewRequest("DELETE", endpoint, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationResponse

	response, err := piper.client.Do(ctx, request, &record)
	if err != nil {
		return nil, response, err
	}

	return record, response, nil
}
