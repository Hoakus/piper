package pipedrive

import (
	"github.com/joho/godotenv"
	"testing"
)

func TestNewClient(t *testing.T) {
	dotenv := godotenv.Load()
	dotenv.
		newClient := pipedrive
}
