package pipedrive

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const customTimeLayout = "2006-01-02 15:04:05"
const standardTimeLayout = "2006-01-02T15:04:05Z07:00"

type AdditionalData struct {
	Pagination struct {
		Start                 int  `json:"start"`
		Limit                 int  `json:"limit"`
		MoreItemsInCollection bool `json:"more_items_in_collection"`
		NextStart             int  `json:"next_start"`
	} `json:"pagination"`
}

type Monetary struct {
	Amount   int    `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

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

// The dates returned by pipedrive has inconsistent formatting depending
// on the module and api version. so we check both before erroring
func (ct *TimeStamp) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}

	var t time.Time

	if strings.ContainsAny(s, "TZ") {
		t, err = time.Parse(customTimeLayout, s)
	} else {
		t, err = time.Parse(standardTimeLayout, s)
	}

	if err != nil {
		return fmt.Errorf("Failed to parse Timestamp during UnmarshalJSON: %v", s)
	}

	ct.Time = t
	return nil
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
