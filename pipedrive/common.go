package pipedrive

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const customTimeLayout = "2006-01-02 15:04:05"
const standardTimeLayout = "2006-01-02T15:04:05Z07:00"

type TimeStamp struct {
	time.Time
}

func (ts TimeStamp) String() string {
	return ts.Time.String()
}

func (ts TimeStamp) DateString() string {
	return ts.Time.Format("2000-11-30")
}

func (ts TimeStamp) DateTimeString() string {
	return ts.Time.Format("2000-11-30 22:10:55")
}
func (ct *TimeStamp) UnmarshalJSON(b []byte) (err error) {
	// Trim the quotes from the JSON string value
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	// Parse the string using the custom layout
	t, err := time.Parse(customTimeLayout, s)
	if err != nil {
		// if that fails, try the standard layout
		t, err = time.Parse(standardTimeLayout, s)
		if err != nil {
			// return an error if neither worked
			return fmt.Errorf("Failed to parse time with either methods: %v", s)
		}
	}
	ct.Time = t
	return nil
}

// Additional common responses are defined here
// this is to be able to unmarshal the pipedrive response
// into more bite-sized structs, in order to limit the amount of overwhelming
// options when extracting data from the *(piper)Response as a user
type AdditionalData struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Start                 int  `json:"start"`
	Limit                 int  `json:"limit"`
	MoreItemsInCollection bool `json:"more_items_in_collection"`
	NextStart             int  `json:"next_start"`
}

func Stringify(obj any) string {
	if obj == nil {
		return "<nil>"
	}
	switch v := obj.(type) {
	case *TimeStamp:
		return v.String()
	default:
		data, err := json.Marshal(obj)
		if err == nil {
			return string(data)
		}
		return fmt.Sprintf("%#v", obj)
	}
}
