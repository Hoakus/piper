package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

type TaskPiper piper

type TasksResponse struct {
	Success        bool            `json:"success"`
	Tasks          []Task          `json:"data"`
	AdditionalData *AdditionalData `json:"additional_data"`
}

type TaskResponse struct {
	Success bool `json:"success"`
	Tasks   Task `json:"data"`
}

// https://developers.pipedrive.com/docs/api/v1/Tasks#getTasks
func (piper *TaskPiper) GetAll(ctx context.Context, params TasksGetOptions) (*TasksResponse, *http.Response, error) {
	endpoint := "tasks"

	request, err := piper.client.NewRequest("GET", endpoint, "1", &params, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *TasksResponse

	response, err := piper.client.Do(ctx, request, &record)
	if err != nil {
		return nil, nil, err
	}

	return record, response, nil
}

// https://developers.pipedrive.com/docs/api/v1/Tasks#getTask
func (piper *TaskPiper) Get(ctx context.Context, taskID int) (*TaskResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("tasks/%d", taskID)

	request, err := piper.client.NewRequest("GET", endpoint, "1", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *TaskResponse

	response, err := piper.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, nil
}

// https://developers.pipedrive.com/docs/api/v1/Tasks#addTask
func (piper *TaskPiper) Add(ctx context.Context, title string, projectID int, body TaskAddOptions) (*TaskResponse, *http.Response, error) {
	endpoint := "tasks"

	// ensuring required fields are set
	body.Title = title
	body.ProjectID = projectID

	request, err := piper.client.NewRequest("POST", endpoint, "1", nil, body)

	if err != nil {
		return nil, nil, err
	}

	var record *TaskResponse

	response, err := piper.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, err
}

// https://developers.pipedrive.com/docs/api/v1/Tasks#updateTask
func (piper *TaskPiper) Update(ctx context.Context, taskID int, params TaskUpdateOptions) (*TaskResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("tasks/%d", taskID)

	request, err := piper.client.NewRequest("PUT", endpoint, "1", params, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *TaskResponse

	response, err := piper.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, err
}

// https://developers.pipedrive.com/docs/api/v1/Tasks#deleteTask
// marks a task as deleted - if it has subtasks, they will also be deleted
func (piper *TaskPiper) Delete(ctx context.Context, taskID int) (*TaskResponse, *http.Response, error) {
	endpoint := fmt.Sprintf("tasks/%d", taskID)

	request, err := piper.client.NewRequest("DELETE", endpoint, "1", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *TaskResponse

	response, err := piper.client.Do(ctx, request, &record)

	if err != nil {
		return nil, nil, err
	}

	return record, response, err
}
