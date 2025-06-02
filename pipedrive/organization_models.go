package pipedrive

import (
	"time"
)

// Defines the different type of Organization responses
// when accessing certain endpoints. When you call a piper.Method,
// on of these 'records' is returned along with the HTTPResponse

type Organization struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	AddTime       time.Time `json:"add_time"`
	UpdateTime    time.Time `json:"update_time"`
	VisibleTo     int       `json:"visible_to"`
	OwnerID       int       `json:"owner_id"`
	LabelIds      []any     `json:"label_ids"`
	Website       any       `json:"website"`
	Linkedin      any       `json:"linkedin"`
	Industry      any       `json:"industry"`
	AnnualRevenue any       `json:"annual_revenue"`
	EmployeeCount any       `json:"employee_count"`
	IsDeleted     bool      `json:"is_deleted"`
	Address       struct {
		Value            string `json:"value"`
		StreetNumber     any    `json:"street_number"`
		Route            any    `json:"route"`
		Sublocality      any    `json:"sublocality"`
		Locality         any    `json:"locality"`
		AdminAreaLevel1  any    `json:"admin_area_level_1"`
		AdminAreaLevel2  any    `json:"admin_area_level_2"`
		Country          any    `json:"country"`
		PostalCode       any    `json:"postal_code"`
		FormattedAddress any    `json:"formatted_address"`
	} `json:"address"`
	NextActivityID          int `json:"next_activity_id,omitempty"`
	LastActivityID          int `json:"last_activity_id,omitempty"`
	OpenDealsCount          int `json:"open_deals_count,omitempty"`
	RelatedOpenDealsCount   int `json:"related_open_deals_count,omitempty"`
	ClosedDealsCount        int `json:"closed_deals_count,omitempty"`
	RelatedClosedDealsCount int `json:"related_closed_deals_count,omitempty"`
	EmailMessagesCount      int `json:"email_messages_count,omitempty"`
	PeopleCount             int `json:"people_count,omitempty"`
	ActivitiesCount         int `json:"activities_count,omitempty"`
	DoneActivitiesCount     int `json:"done_activities_count,omitempty"`
	UndoneActivitiesCount   int `json:"undone_activities_count,omitempty"`
	FilesCount              int `json:"files_count,omitempty"`
	NotesCount              int `json:"notes_count,omitempty"`
	FollowersCount          int `json:"followers_count,omitempty"`
	WonDealsCount           int `json:"won_deals_count,omitempty"`
	RelatedWonDealsCount    int `json:"related_won_deals_count,omitempty"`
	LostDealsCount          int `json:"lost_deals_count,omitempty"`
	RelatedLostDealsCount   int `json:"related_lost_deals_count,omitempty"`
	//Custom_Fields
	CustomFields any `json:"custom_fields,omitempty"`
}
