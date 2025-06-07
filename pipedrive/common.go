package pipedrive

import (
	"encoding/json"
	"fmt"
	"time"
)

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
