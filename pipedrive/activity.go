package pipedrive

import (
	"context"
	"fmt"
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

func (piper *ActivitiesPiper) Get(ctx context.Context, activityID int, params GetActivityOpts) (*ActivityResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("api/v2/activities/%d", activityID)

	request, err := piper.client.NewRequest("GET", endpoint, params, nil)

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

// https://developers.pipedrive.com/docs/api/v1/Activities#addActivity
func (piper *ActivitiesPiper) Add(ctx context.Context, body AddActivityOpts) (*ActivityResponse, *http.Response, error) {
	endpoint := "api/v2/activities"

	request, err := piper.client.NewRequest("POST", endpoint, nil, body)

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

// https://developers.pipedrive.com/docs/api/v1/Activities#updateActivity
func (piper *ActivitiesPiper) Update(ctx context.Context, activityID int, body UpdateActivityOpts) (*ActivityResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("api/v2/activities/%d", activityID)

	request, err := piper.client.NewRequest("PATCH", endpoint, nil, body)

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

// https://developers.pipedrive.com/docs/api/v1/Activities#deleteActivity
func (piper *ActivitiesPiper) Delete(ctx context.Context, activityID int) (*ActivityResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("api/v2/activities/%d", activityID)

	request, err := piper.client.NewRequest("DELETE", endpoint, nil, nil)

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
