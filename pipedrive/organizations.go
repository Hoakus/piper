package pipedrive

import (
	"context"
	"net/http"
	"strconv"
)

type OrganizationPiper piper

type OrganizationResponse struct {
	Success        bool            `json:"success"`
	Data           Organization    `json:"data"`
	AdditionalData *AdditionalData `json:"additional_data,omitempty"`
	RelatedObjects *RelatedObjects `json:"related_objects,omitempty"`
}

type OrganizationsResponse struct {
	Success        bool            `json:"success"`
	Data           []Organization  `json:"data"`
	AdditionalData *AdditionalData `json:"additional_data,omitempty"`
	RelatedObjects *RelatedObjects `json:"related_objects,omitempty"`
}

// https://developers.pipedrive.com/docs/api/v1/Organizations#getOrganizations
func (piper *OrganizationPiper) GetAll(ctx context.Context, params OrganizationGetAllOpts) (*OrganizationsResponse, *http.Response, error) {
	endpoint := "organizations"

	request, err := piper.client.NewRequest("GET", endpoint, "2", &params, nil)

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
func (piper *OrganizationPiper) Get(ctx context.Context, record_id int, params OrganizationGetOpts) (*OrganizationResponse, *http.Response, error) {
	endpoint := "organizations" + "/" + strconv.Itoa(record_id)

	request, err := piper.client.NewRequest("GET", endpoint, "2", params, nil)

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
func (piper *OrganizationPiper) Add(ctx context.Context, body OrganizationAddOpts) (*OrganizationResponse, *http.Response, error) {
	endpoint := "organizations"

	request, err := piper.client.NewRequest("POST", endpoint, "2", nil, body)

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
func (piper *OrganizationPiper) Update(ctx context.Context, record_id int, body OrganizationUpdateOpts) (*OrganizationResponse, *http.Response, error) {
	endpoint := "organizations" + "/" + strconv.Itoa(record_id)

	request, err := piper.client.NewRequest("PATCH", endpoint, "2", nil, body)

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
func (piper *OrganizationPiper) Delete(ctx context.Context, record_id int) (*OrganizationResponse, *http.Response, error) {
	endpoint := "organizations" + "/" + strconv.Itoa(record_id)

	request, err := piper.client.NewRequest("DELETE", endpoint, "2", nil, nil)

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
