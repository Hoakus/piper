package pipedrive

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
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
