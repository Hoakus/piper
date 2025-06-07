# PIPER

Pipedrive API package written in Go

## Installation

```go get github.com/Hoakus/piper/pipedrive```

## Usage


```
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

	// see the resulting struct record
	val, _ := json.MarshalIndent(record, "", "\t")
	fmt.Println(string(val))

}```

## Notes

**Under Regular Development**
Expect breaking changes if you use this package. This is my first undertaking of this size,
new to golang, and in general I have little to zero idea what I am doing. If you would like
to contribute to a project that actually has license, and one that is going to be continuosly
updated to ensure compatibility with current pipedrive api, SHOW ME WHAT YOU GOT. 

Thanks!

