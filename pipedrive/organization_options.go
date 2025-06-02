package pipedrive

import (
	"encoding/json"
	"fmt"
)

type OrganizationGetAllOpts struct {
	FilterID      int      `url:"filter_id,omitempty"`
	IDs           []string `url:"ids,omitempty"`
	OwnerID       int      `url:"owner_id,omitempty"`
	UpdatedSince  string   `url:"updated_since,omitempty"`
	UpdatedUntil  string   `url:"updated_until,omitempty"`
	SortBy        string   `url:"sort_by,omitempty"`
	SortDirection string   `url:"sort_direction,omitempty"`
	IncludeFields string   `url:"include_fields,omitempty"` // comma-separatedlist
	CustomFields  string   `url:"custom_fields,omitempty"`  // comma-separated list
	Limit         int      `url:"limit,omitempty"`
	Cursor        string   `url:"cursor,omitempty"`
}

func (p OrganizationGetAllOpts) String() string {
	return Stringify(p)
}

type OrganizationGetOpts struct {
	IncludeFields []string `url:"include_fields,omitempty"`
	CustomFields  []string `url:"custom_fields,omitempty"`
}

func (p OrganizationGetOpts) String() string {
	return Stringify(p)
}

type OrganizationAddOpts struct {
	Name         string     `json:"name"`
	OwnerID      int        `json:"owner_id,omitempty"`
	AddTime      *TimeStamp `json:"add_time,omitempty"`
	UpdateTime   *TimeStamp `json:"update_time,omitempty"`
	VisibleTo    int        `json:"visible_to,omitempty"`
	LabelIDs     []int      `json:"label_ids,omitempty"`
	CustomFields []string   `json:"custom_fields,omitempty"`
}

func (o OrganizationAddOpts) String() string {
	return Stringify(o)
}

type OrganizationUpdateOpts struct {
	Name         string     `json:"name"`
	OwnerID      int        `json:"owner_id,omitempty"`
	AddTime      *TimeStamp `json:"add_time,omitempty"`
	UpdateTime   *TimeStamp `json:"update_time,omitempty"`
	VisibleTo    int        `json:"visible_to,omitempty"`
	LabelIDs     []int      `json:"label_ids,omitempty"`
	CustomFields []string   `json:"custom_fields,omitempty"`
}

func (o OrganizationUpdateOpts) String() string {
	return Stringify(o)
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
