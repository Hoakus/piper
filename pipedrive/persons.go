package pipedrive

import (
	"context"
	"net/http"
)

type PersonsPiper piper

type PersonsResponse struct {
	Success        bool           `json:"success"`
	Data           []Person       `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type PersonResponse struct {
	Success bool   `json:"success"`
	Data    Person `json:"data"`
}

func (p *PersonsPiper) Add(ctx context.Context, body AddPersonOpts) (*PersonResponse, *http.Response, error) {
	endpoint := "api/v2/persons"

	req, err := p.client.NewRequest("POST", endpoint, nil, body)
	if err != nil {
		return nil, nil, err
	}

	var record *PersonResponse

	res, err := p.client.Do(ctx, req, &record)
	if err != nil {
		return nil, nil, err
	}

	return record, res, nil
}
