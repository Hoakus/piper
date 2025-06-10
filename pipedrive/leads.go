package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

type LeadsPiper piper

type LeadsResponse struct {
	Success        bool           `json:"success"`
	Data           []Lead         `json:"data"`
	AdditionalData AdditionalData `json:"addtional_data"`
}

type LeadResponse struct {
	Success bool `json:"success"`
	Data    Lead `json:"data"`
}

// https://developers.pipedrive.com/docs/api/v1/Leads#getLead
func (piper *LeadsPiper) Get(ctx context.Context, leadID string) (*LeadResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("v1/leads/%s", leadID)

	params := leadsParams{APIKey: piper.client.apiKey}

	request, err := piper.client.NewRequest("GET", endpoint, params, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *LeadResponse

	response, err := piper.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, nil
}

// https://developers.pipedrive.com/docs/api/v1/Leads#addLead
// must provide either an organization ID or person ID in options
func (piper *LeadsPiper) Add(ctx context.Context, body AddLeadOpts) (*LeadResponse, *http.Response, error) {
	endpoint := "v1/leads"

	params := leadsParams{APIKey: piper.client.apiKey}

	request, err := piper.client.NewRequest("POST", endpoint, params, body)

	if err != nil {
		return nil, nil, err
	}

	var record *LeadResponse

	response, err := piper.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, nil
}

// https://developers.pipedrive.com/docs/api/v1/Leads#updateLead
func (piper *LeadsPiper) Update(ctx context.Context, leadID string) (*LeadResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("v1/leads/%s", leadID)

	params := leadsParams{APIKey: piper.client.apiKey}

	request, err := piper.client.NewRequest("PATCH", endpoint, params, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *LeadResponse

	response, err := piper.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, nil
}

// https://developers.pipedrive.com/docs/api/v1/Leads#deleteLead
func (piper *LeadsPiper) Delete(ctx context.Context, leadID string) (*LeadResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("v1/leads/%s", leadID)

	request, err := piper.client.NewRequest("DELETE", endpoint, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *LeadResponse

	response, err := piper.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, nil
}
