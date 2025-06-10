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

	ctx := context.Background()

	showDealFields(client, ctx)

}

func showDealFields(c *pipedrive.Client, ctx context.Context) {
	record, _, err := c.DealFields.GetAll(ctx, pipedrive.GetDealsFieldsOpts{})

	if err != nil {
		log.Fatalf("Failed to execute DealFields.GetAll(): %v", err)
	}

	for i, item := range record.Data {
		fmt.Println("===============================")
		fmt.Println("Item number:", i)
		fmt.Println("Field Name:", item.Name)
		fmt.Println("Field Key:", item.Key)
		fmt.Println("Field Type:", item.FieldType)
		fmt.Println("Is Mandatory:", item.MandatoryFlag)
		fmt.Println("Options:", item.Options)
		fmt.Println("===============================")
	}

	PrintResult(record)

}

func PrintResult(record any) {
	val, _ := json.Marshal(record)

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
