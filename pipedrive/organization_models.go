package pipedrive

type Organization struct {
	ID                      int        `json:"id"`
	Name                    string     `json:"name"`
	AddTime                 *TimeStamp `json:"add_time"`
	UpdateTime              *TimeStamp `json:"update_time"`
	VisibleTo               int        `json:"visible_to"`
	OwnerID                 int        `json:"owner_id"`
	LabelIds                []any      `json:"label_ids"`
	Website                 string     `json:"website,omitempty"`
	Linkedin                string     `json:"linkedin,omitempty"`
	Industry                int        `json:"industry"`
	AnnualRevenue           int        `json:"annual_revenue"`
	EmployeeCount           int        `json:"employee_count"`
	IsDeleted               bool       `json:"is_deleted"`
	Address                 *Address   `json:"address,omitempty"`
	NextActivityID          int        `json:"next_activity_id,omitempty"`
	LastActivityID          int        `json:"last_activity_id,omitempty"`
	OpenDealsCount          int        `json:"open_deals_count,omitempty"`
	RelatedOpenDealsCount   int        `json:"related_open_deals_count,omitempty"`
	ClosedDealsCount        int        `json:"closed_deals_count,omitempty"`
	RelatedClosedDealsCount int        `json:"related_closed_deals_count,omitempty"`
	EmailMessagesCount      int        `json:"email_messages_count,omitempty"`
	PeopleCount             int        `json:"people_count,omitempty"`
	ActivitiesCount         int        `json:"activities_count,omitempty"`
	DoneActivitiesCount     int        `json:"done_activities_count,omitempty"`
	UndoneActivitiesCount   int        `json:"undone_activities_count,omitempty"`
	FilesCount              int        `json:"files_count,omitempty"`
	NotesCount              int        `json:"notes_count,omitempty"`
	FollowersCount          int        `json:"followers_count,omitempty"`
	WonDealsCount           int        `json:"won_deals_count,omitempty"`
	RelatedWonDealsCount    int        `json:"related_won_deals_count,omitempty"`
	LostDealsCount          int        `json:"lost_deals_count,omitempty"`
	RelatedLostDealsCount   int        `json:"related_lost_deals_count,omitempty"`
	CustomFields            any        `json:"custom_fields,omitempty"`
}

type Address struct {
	Value            string `json:"value,omitempty"`
	StreetNumber     any    `json:"street_number,omitempty"`
	Route            any    `json:"route,omitempty"`
	Sublocality      any    `json:"sublocality,omitempty"`
	Locality         any    `json:"locality,omitempty"`
	AdminAreaLevel1  any    `json:"admin_area_level_1,omitempty"`
	AdminAreaLevel2  any    `json:"admin_area_level_2,omitempty"`
	Country          any    `json:"country,omitempty"`
	PostalCode       any    `json:"postal_code,omitempty"`
	FormattedAddress any    `json:"formatted_address,omitempty"`
}

type GetOrganizationsOpts struct {
	FilterID      int      `url:"filter_id,omitempty"`
	IDs           []string `url:"ids,omitempty"`
	OwnerID       int      `url:"owner_id,omitempty"`
	UpdatedSince  string   `url:"updated_since,omitempty"`
	UpdatedUntil  string   `url:"updated_until,omitempty"`
	SortBy        string   `url:"sort_by,omitempty"`
	SortDirection string   `url:"sort_direction,omitempty"`
	Limit         int      `url:"limit,omitempty"`
	Cursor        string   `url:"cursor,omitempty"`
	IncludeFields string   `url:"include_fields,omitempty"` // comma-separated string
	CustomFields  any      `url:"custom_fields,omitempty"`  // comma-separated string
}

func (p GetOrganizationsOpts) String() string {
	return Stringify(p)
}

type GetOrganizationOpts struct {
	IncludeFields string `url:"include_fields,omitempty"` // comma-separated string
	CustomFields  string `url:"custom_fields,omitempty"`  // comma-separated string
}

func (p GetOrganizationOpts) String() string {
	return Stringify(p)
}

type AddOrganizationOpts struct {
	Name         string     `json:"name"` // required field
	OwnerID      int        `json:"owner_id,omitempty"`
	AddTime      *TimeStamp `json:"add_time,omitempty"`
	UpdateTime   *TimeStamp `json:"update_time,omitempty"`
	VisibleTo    int        `json:"visible_to,omitempty"`
	LabelIDs     []int      `json:"label_ids,omitempty"`
	Address      *Address   `json:"address,omitempty"`
	CustomFields any        `json:"custom_fields,omitempty"`
}

func (p AddOrganizationOpts) String() string {
	return Stringify(p)
}

// CustomFields should be a struct where the field name is the pipedrive api
// key for that field, and the type corresponding to the field type
// you have set in your pipedrive instance
type UpdateOrganizationOpts struct {
	Name         string     `json:"name,omitempty"`
	OwnerID      int        `json:"owner_id,omitempty"`
	AddTime      *TimeStamp `json:"add_time,omitempty"`
	UpdateTime   *TimeStamp `json:"update_time,omitempty"`
	VisibleTo    int        `json:"visible_to,omitempty"`
	LabelIDs     []int      `json:"label_ids,omitempty"`
	Address      *Address   `json:"address,omitempty"`
	CustomFields any        `json:"custom_fields,omitempty"`
}

func (o UpdateOrganizationOpts) String() string {
	return Stringify(o)
}
