package pipedrive

import (
	"context"
	"net/http"
)

type ActivitiesPiper piper

type ActivitiesResponse struct {
	Success        bool           `json:"success"`
	Data           []Activity     `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type ActivityResponse struct {
	Success bool     `json:"success"`
	Data    Activity `json:"data"`
}

func (piper *ActivitiesPiper) Add(ctx context.Context, params ActivitiesAddOptions) (*ActivityResponse, *http.Response, error) {
	endpoint := "api/v2/activities"

	request, err := piper.client.NewRequest("POST", endpoint, params, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *ActivityResponse

	response, err := piper.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, err
}
