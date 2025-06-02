package pipedrive

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Response *http.Response
	Details  ErrorDetails
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Failed to %s with %s : %v",
		e.Response.Request.Method, e.Response.Request.URL.String(),
		e.Details.String())
}

type ErrorDetails struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode int    `json:"errorCode"`
	ErrorInfo string `json:"error_info"`
}

func (e ErrorDetails) String() string {
	return Stringify(e)
}
