package pipedrive

import (
	"time"
)

// Defines the different type of Organization responses
// when accessing certain endpoints. When you call a piper.Method,
// on of these 'records' is returned along with the HTTPResponse
type Organization struct {
	ID      int `json:"id"`
	OwnerID struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		HasPic     int    `json:"has_pic"`
		PicHash    any    `json:"pic_hash"`
		ActiveFlag bool   `json:"active_flag"`
		Value      int    `json:"value"`
	} `json:"owner_id"`
	Name                    string `json:"name"`
	OpenDealsCount          int    `json:"open_deals_count"`
	RelatedOpenDealsCount   int    `json:"related_open_deals_count"`
	ClosedDealsCount        int    `json:"closed_deals_count"`
	RelatedClosedDealsCount int    `json:"related_closed_deals_count"`
	EmailMessagesCount      int    `json:"email_messages_count"`
	PeopleCount             int    `json:"people_count"`
	ActivitiesCount         int    `json:"activities_count"`
	DoneActivitiesCount     int    `json:"done_activities_count"`
	UndoneActivitiesCount   int    `json:"undone_activities_count"`
	FilesCount              int    `json:"files_count"`
	NotesCount              int    `json:"notes_count"`
	FollowersCount          int    `json:"followers_count"`
	WonDealsCount           int    `json:"won_deals_count"`
	RelatedWonDealsCount    int    `json:"related_won_deals_count"`
	LostDealsCount          int    `json:"lost_deals_count"`
	RelatedLostDealsCount   int    `json:"related_lost_deals_count"`
	ActiveFlag              bool   `json:"active_flag"`
	PictureID               any    `json:"picture_id"`
	CountryCode             any    `json:"country_code"`
	FirstChar               string `json:"first_char"`
	UpdateTime              string `json:"update_time"`
	DeleteTime              any    `json:"delete_time"`
	AddTime                 string `json:"add_time"`
	VisibleTo               string `json:"visible_to"`
	NextActivityDate        any    `json:"next_activity_date"`
	NextActivityTime        any    `json:"next_activity_time"`
	NextActivityID          any    `json:"next_activity_id"`
	LastActivityID          any    `json:"last_activity_id"`
	LastActivityDate        any    `json:"last_activity_date"`
	Label                   any    `json:"label"`
	LabelIds                []any  `json:"label_ids"`
	Address                 string `json:"address"`
	AddressSubpremise       any    `json:"address_subpremise"`
	AddressStreetNumber     any    `json:"address_street_number"`
	AddressRoute            any    `json:"address_route"`
	AddressSublocality      any    `json:"address_sublocality"`
	AddressLocality         any    `json:"address_locality"`
	AddressAdminAreaLevel1  any    `json:"address_admin_area_level_1"`
	AddressAdminAreaLevel2  any    `json:"address_admin_area_level_2"`
	AddressCountry          any    `json:"address_country"`
	AddressPostalCode       any    `json:"address_postal_code"`
	AddressFormattedAddress any    `json:"address_formatted_address"`
	Website                 any    `json:"website"`
	Linkedin                any    `json:"linkedin"`
	Industry                any    `json:"industry"`
	AnnualRevenue           any    `json:"annual_revenue"`
	EmployeeCount           any    `json:"employee_count"`
	OwnerName               string `json:"owner_name"`
	CcEmail                 string `json:"cc_email"`
	CompanyID               int    `json:"company_id"`
	// Define your fields
}

type OrganizationDetails struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	AddTime       time.Time `json:"add_time"`
	UpdateTime    time.Time `json:"update_time"`
	VisibleTo     int       `json:"visible_to"`
	CustomFields  struct{}  `json:"custom_fields"` // Define your fields
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
}
