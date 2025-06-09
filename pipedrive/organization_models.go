package pipedrive

// Defines the different type of Organization responses
// when accessing certain endpoints. When you call a piper.Method,
// on of these 'records' is returned along with the HTTPResponse
type Organization struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	AddTime       *TimeStamp `json:"add_time"`
	UpdateTime    *TimeStamp `json:"update_time"`
	VisibleTo     int        `json:"visible_to"`
	OwnerID       int        `json:"owner_id"`
	LabelIds      []any      `json:"label_ids"`
	Website       string     `json:"website,omitempty"`
	Linkedin      string     `json:"linkedin,omitempty"`
	Industry      int        `json:"industry"`
	AnnualRevenue int        `json:"annual_revenue"`
	EmployeeCount int        `json:"employee_count"`
	IsDeleted     bool       `json:"is_deleted"`
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
	CustomFields            any `json:"custom_fields,omitempty"`
}

// IncludeFields and CustomFields must be comma-separated string with each
// substring being the field name you wish to be included in the response
type OrganizationsGetOptions struct {
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

func (p OrganizationsGetOptions) String() string {
	return Stringify(p)
}

// IncludeFields and CustomFields must be comma-separated string with each
// substring being the field name you wish to be included in the response
type OrganizationGetOptions struct {
	IncludeFields string `url:"include_fields,omitempty"`
	CustomFields  string `url:"custom_fields,omitempty"`
}

func (p OrganizationGetOptions) String() string {
	return Stringify(p)
}

// CustomFields should be a struct where the field name is the pipedrive api
// key for that field, and the type corresponding to the field type
// you have set in your pipedrive instance
type OrganizationAddOptions struct {
	Name         string     `json:"name"`
	OwnerID      int        `json:"owner_id,omitempty"`
	AddTime      *TimeStamp `json:"add_time,omitempty"`
	UpdateTime   *TimeStamp `json:"update_time,omitempty"`
	VisibleTo    int        `json:"visible_to,omitempty"`
	LabelIDs     []int      `json:"label_ids,omitempty"`
	CustomFields any        `json:"custom_fields,omitempty"`
}

func (o OrganizationAddOptions) String() string {
	return Stringify(o)
}

// CustomFields should be a struct where the field name is the pipedrive api
// key for that field, and the type corresponding to the field type
// you have set in your pipedrive instance
type OrganizationUpdateOptions struct {
	Name         string     `json:"name"`
	OwnerID      int        `json:"owner_id,omitempty"`
	AddTime      *TimeStamp `json:"add_time,omitempty"`
	UpdateTime   *TimeStamp `json:"update_time,omitempty"`
	VisibleTo    int        `json:"visible_to,omitempty"`
	LabelIDs     []int      `json:"label_ids,omitempty"`
	CustomFields any        `json:"custom_fields,omitempty"`
}

func (o OrganizationUpdateOptions) String() string {
	return Stringify(o)
}
