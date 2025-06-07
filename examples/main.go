package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Hoakus/piper/pipedrive"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load your secrets
	APIKey := getCredentials("API_KEY")
	DomainName := getCredentials("DOMAIN_NAME")

	// create configuration
	cfg := pipedrive.Config{APIKey: APIKey, DomainName: DomainName}

	// declare client and pass in ptr to your configuration
	client := pipedrive.NewClient(&cfg)

	// pipedrive modules are called "pipers" internally
	// access any piper by it's module name - i.e "Organization"
	// followed by an endpoint method
	record, _, err := client.Organization.Get(context.Background(), 2366, pipedrive.OrganizationGetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	val, _ := json.MarshalIndent(record, "", "\t")

	// see the resulting struct record
	fmt.Println(string(val))

}

func getCredentials(env_key string) string {
	godotenv.Load()
	env_var := os.Getenv(env_key)

	if env_var == "" {
		log.Fatalf("could not load credentials '%v'", env_var)
	}

	return env_var
}
