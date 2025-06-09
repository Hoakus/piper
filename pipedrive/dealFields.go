package pipedrive

import (
	"context"
	"net/http"
)

type DealFieldsPiper piper

type DealsFieldsResponse struct {
	Success        bool           `json:"success"`
	Data           []DealFields   `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type DealFieldsResponse struct {
	Success bool       `json:"success"`
	Data    DealFields `json:"data"`
}

// https://developers.pipedrive.com/docs/api/v1/DealFields#getDealFields
func (d *DealFieldsPiper) GetAll(ctx context.Context, params GetDealsFieldsOptions) (*DealsFieldsResponse, *http.Response, error) {
	endpoint := "api/v1/dealFields"

	request, err := d.client.NewRequest("GET", endpoint, params, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealsFieldsResponse

	response, err := d.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, nil
}
