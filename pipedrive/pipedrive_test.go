package pipedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func getTestingCredentials(key string) string {
	godotenv.Load()

	env_var := os.Getenv(key)
	if env_var == "" {
		log.Fatalf("Could not load enviroment variable %v", key)
	}

	return env_var

}

func getTestClient() *Client {
	APIKey := getTestingCredentials("API_KEY")
	DomainName := getTestingCredentials("DOMAIN_NAME")

	cfg := &Config{
		APIKey:     APIKey,
		DomainName: DomainName,
	}

	return NewClient(cfg)

}

func TestGetAll(t *testing.T) {
	client := getTestClient()

	testCases := []struct {
		name           string
		ctx            context.Context
		params         OrganizationGetAllOpts
		expectedOutput string
		expectError    bool
	}{
		{
			name: "Test correct use",
			ctx:  context.TODO(),
			params: OrganizationGetAllOpts{
				IncludeFields: "closed_deals_count,notes_count",
				Limit:         3,
			},
			expectedOutput: "?include_fields=closed_deals_count%2cnotes_count&limit=3",
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, response, err := client.Organization.GetAll(tc.ctx, tc.params)
			if err != nil {
				t.Fatalf("%s failed. recieved error : %v",
					tc.name, err)
			}

			actualOutput := response.Request.URL.String()

			if !strings.Contains(actualOutput, tc.expectedOutput) {
				t.Fatalf("%s failed. expected %v, received %v",
					tc.name, tc.expectedOutput, actualOutput)
			}

			v, _ := json.MarshalIndent(result, "", "\t")

			t.Log(string(v))
		})
	}
}

func TestStringify(t *testing.T) {
	globalTime := time.Now()

	testCases := []struct {
		name           string
		testStruct     fmt.Stringer
		expectedOutput string
		expectError    bool
	}{
		{
			name:           "Test TimeStamp String()",
			testStruct:     TimeStamp{globalTime},
			expectedOutput: globalTime.String(),
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := tc.testStruct.String()

			if !tc.expectError {
				if tc.expectedOutput != actualOutput {
					t.Errorf("failed %v : expected %v but received %v",
						tc.name, tc.expectedOutput, actualOutput)
				}
			} else {
				if tc.expectedOutput == actualOutput {
					t.Errorf("failed %v : expected %v but received %v",
						tc.name, tc.expectedOutput, actualOutput)
				}

			}

		})

	}

}

func TestNewRequest(t *testing.T) {
	client := getTestClient()

	testCases := []struct {
		name           string
		method         string
		endpoint       string
		apiversion     string
		options        any
		body           any
		expectedOutput string
		expectError    bool
	}{
		// --- Test Case 1: Happy Path ---
		{
			name:       "Test Valid Request",
			method:     "GET",
			endpoint:   "organizations",
			apiversion: "2",
			options: struct {
				ID    int
				Limit int
			}{
				ID:    2255,
				Limit: 10,
			},
			body:           nil,
			expectedOutput: fmt.Sprintf("%vapi/v2/organizations?id=2255&limit=10", client.BaseURL),
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := client.NewRequest(
				tc.method,
				tc.endpoint,
				tc.apiversion,
				tc.options,
				tc.body,
			)

			actualOutput := req.URL

			if tc.expectError {
				if err == nil {
					t.Errorf("expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error, but got: %v", err)
				}
			}

			compareURLQueries(t, tc.expectedOutput, actualOutput.String())

		})
	}
}

func compareURLQueries(t *testing.T, expectedURLStr, actualURLStr string) {
	expectedURL, err := url.Parse(expectedURLStr)
	if err != nil {
		t.Fatalf("Failed to parse expected URL: %v", err)
	}
	actualURL, err := url.Parse(actualURLStr)
	if err != nil {
		t.Fatalf("Failed to parse actual URL: %v", err)
	}

	if expectedURL.Scheme != actualURL.Scheme {
		t.Errorf("Scheme mismatch: expected %s, got %s", expectedURL.Scheme, actualURL.Scheme)
	}
	if expectedURL.Host != actualURL.Host {
		t.Errorf("Host mismatch: expected %s, got %s", expectedURL.Host, actualURL.Host)
	}
	if expectedURL.Path != actualURL.Path {
		t.Errorf("Path mismatch: expected %s, got %s", expectedURL.Path, actualURL.Path)
	}

	expectedQuery := expectedURL.Query()
	actualQuery := actualURL.Query()

	if !reflect.DeepEqual(expectedQuery, actualQuery) {
		t.Errorf("Query mismatch:\nExpected: %v\nActual:   %v", expectedQuery, actualQuery)
	} else {

		t.Logf("All tests successful %v", actualURL.String())
	}
}
