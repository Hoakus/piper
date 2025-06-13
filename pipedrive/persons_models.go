package pipedrive

type MarketingStatus int

const (
	NoConsent MarketingStatus = iota
	Unsubscribed
	Subscribed
	Archived
)

func (ms MarketingStatus) String() string {
	s := ""

	switch ms {
	case NoConsent:
		s = "no_consent"
	case Unsubscribed:
		s = "unsubscribed"
	case Subscribed:
		s = "subscribed"
	case Archived:
		s = "archived"
	default:
		s = ""
	}
	return s
}

type Person struct {
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
	OrgID                       any    `json:"org_id"`
	Name                        string `json:"name"`
	FirstName                   string `json:"first_name"`
	LastName                    string `json:"last_name"`
	OpenDealsCount              int    `json:"open_deals_count"`
	RelatedOpenDealsCount       int    `json:"related_open_deals_count"`
	ClosedDealsCount            int    `json:"closed_deals_count"`
	RelatedClosedDealsCount     int    `json:"related_closed_deals_count"`
	ParticipantOpenDealsCount   int    `json:"participant_open_deals_count"`
	ParticipantClosedDealsCount int    `json:"participant_closed_deals_count"`
	EmailMessagesCount          int    `json:"email_messages_count"`
	ActivitiesCount             int    `json:"activities_count"`
	DoneActivitiesCount         int    `json:"done_activities_count"`
	UndoneActivitiesCount       int    `json:"undone_activities_count"`
	FilesCount                  int    `json:"files_count"`
	NotesCount                  int    `json:"notes_count"`
	FollowersCount              int    `json:"followers_count"`
	WonDealsCount               int    `json:"won_deals_count"`
	RelatedWonDealsCount        int    `json:"related_won_deals_count"`
	LostDealsCount              int    `json:"lost_deals_count"`
	RelatedLostDealsCount       int    `json:"related_lost_deals_count"`
	ActiveFlag                  bool   `json:"active_flag"`
	Phone                       []struct {
		Value   string `json:"value"`
		Primary bool   `json:"primary"`
	} `json:"phone"`
	Email []struct {
		Label   string `json:"label"`
		Value   string `json:"value"`
		Primary bool   `json:"primary"`
	} `json:"email"`
	FirstChar            string `json:"first_char"`
	UpdateTime           string `json:"update_time"`
	DeleteTime           any    `json:"delete_time"`
	AddTime              string `json:"add_time"`
	VisibleTo            string `json:"visible_to"`
	PictureID            any    `json:"picture_id"`
	NextActivityDate     any    `json:"next_activity_date"`
	NextActivityTime     any    `json:"next_activity_time"`
	NextActivityID       any    `json:"next_activity_id"`
	LastActivityID       any    `json:"last_activity_id"`
	LastActivityDate     any    `json:"last_activity_date"`
	LastIncomingMailTime any    `json:"last_incoming_mail_time"`
	LastOutgoingMailTime any    `json:"last_outgoing_mail_time"`
	Label                any    `json:"label"`
	LabelIds             []any  `json:"label_ids"`
	Im                   []struct {
		Value   string `json:"value"`
		Primary bool   `json:"primary"`
	} `json:"im"`
	PostalAddress                 any    `json:"postal_address"`
	PostalAddressSubpremise       any    `json:"postal_address_subpremise"`
	PostalAddressStreetNumber     any    `json:"postal_address_street_number"`
	PostalAddressRoute            any    `json:"postal_address_route"`
	PostalAddressSublocality      any    `json:"postal_address_sublocality"`
	PostalAddressLocality         any    `json:"postal_address_locality"`
	PostalAddressAdminAreaLevel1  any    `json:"postal_address_admin_area_level_1"`
	PostalAddressAdminAreaLevel2  any    `json:"postal_address_admin_area_level_2"`
	PostalAddressCountry          any    `json:"postal_address_country"`
	PostalAddressPostalCode       any    `json:"postal_address_postal_code"`
	PostalAddressFormattedAddress any    `json:"postal_address_formatted_address"`
	Notes                         any    `json:"notes"`
	Birthday                      any    `json:"birthday"`
	JobTitle                      any    `json:"job_title"`
	OrgName                       any    `json:"org_name"`
	CcEmail                       string `json:"cc_email"`
	PrimaryEmail                  string `json:"primary_email"`
	OwnerName                     string `json:"owner_name"`
	CompanyID                     int    `json:"company_id"`
}

type AddPersonOpts struct {
	Name       string     `json:"name,omitempty"`
	OwnerID    int        `json:"owner_id,omitempty"`
	OrgID      int        `json:"org_id,omitempty"`
	AddTime    *TimeStamp `json:"add_time,omitempty"`
	UpdateTime *TimeStamp `json:"update_time,omitempty"`
	Phone      []struct {
		Value   string `json:"value"`
		Primary bool   `json:"primary"`
	} `json:"phone,omitempty"`
	Email []struct {
		Label   string `json:"label"`
		Value   string `json:"value"`
		Primary bool   `json:"primary"`
	} `json:"email,omitempty"`
	VisibleTo       int
	LabelIds        []any           `json:"label_ids"`
	MarketingStatus MarketingStatus `json:"marketing_status,omitempty"`
}
