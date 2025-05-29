package pipedrive

import (
	"context"
	"strconv"
)

type OrganizationPiper piper

type OrganizationResponse struct {
	Success        bool           `json:"success"`
	Data           Organization   `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
	RelatedObjects RelatedObjects `json:"related_objects,omitempty"`
}

type OrganizationDetailsResponse struct {
	Success        bool                `json:"success"`
	Data           OrganizationDetails `json:"data"`
	AdditionalData AdditionalData      `json:"additional_data,omitempty"`
	RelatedObjects RelatedObjects      `json:"related_objects,omitempty"`
}

// https://developers.pipedrive.com/docs/api/v1/Organizations#getOrganization
func (piper *OrganizationPiper) GetByID(ctx context.Context, id int, params GetByIDParams) (*OrganizationDetailsResponse, *Response, error) {
	endpoint := "organizations" + "/" + strconv.Itoa(id)

	request, err := piper.client.NewRequest("GET", endpoint, "2", params, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationDetailsResponse

	response, err := piper.client.Do(ctx, request, &record)
	if err != nil {
		return nil, response, err
	}

	return record, response, nil
}
