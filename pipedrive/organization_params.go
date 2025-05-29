package pipedrive

import (
	"encoding/json"
	"fmt"
)

// TODO: Decide on solution for params - abstract or direct struct pass

// https://developers.pipedrive.com/docs/api/v1/Organizations#getOrganization
type GetByIDParams struct {
	IncludeFields []string `json:"include_fields"`
	CustomFields  []string `json:"custom_fields"`
}

func (p GetByIDParams) toString() string {
	return Stringify(p)
}

func Stringify(value any) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		fmt.Println("ERROR IN toSTRING")
	}
	return string(bytes)
}
