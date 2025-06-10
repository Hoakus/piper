# PIPER

Pipedrive API client written in Go

**This project is rapidly undergoing changes.** Expect breaking changes if you
use this package. There are no guarantees of reliability at this point in time.
The public API may see significant changes in the coming updates.

## Installation

```go get github.com/Hoakus/piper/pipedrive```

## Usage

``` go
func main() {
    // load your secrets
    apiKey := getCredentials("API_KEY")
    myDomain := getCredentials("DOMAIN_NAME")

    // create configuration
    cfg := pipedrive.Config{APIKey: apiKey, DomainName: myDomain}

    // initialize the pipedrive client with your config
    client := pipedrive.NewClient(&cfg)

    // access any piper by it's module name - i.e "Organization"
    record, _, err := client.Organization.Get(context.Background(), 2366, pipedrive.OrganizationGetOptions{})

    if err != nil {
        log.Fatalf("Could not get Organization record: %v", err)
    }

    // see the resulting struct record
    val, _ := json.MarshalIndent(record, "", "\t")

    fmt.Println(string(val))

}
```

## Notes

### Functionality

There are still several pipedrive modules absent, with current modules containing
only partial implementation of all available endpoints.

There are models provided for each current endpoint method. If you want to use
custom fields, you must define them separately, and add them to the endpoint
options struct provided in the package.

### Rate Limiting

The main client method checks for context signals, and whether the next request
would fail due to current rate limit. If it would fail, an error is returned
for the user to manage their own rate limiting implementation.

The Rate struct holds a mutex to allow for concurrent use.
