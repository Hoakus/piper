package pipedrive

type Lead struct {
	ID                string     `json:"id"`
	Title             string     `json:"title"`
	OwnerID           int        `json:"owner_id"`
	CreatorID         int        `json:"creator_id"`
	LabelIDs          []string   `json:"label_ids"`
	Value             *Monetary  `json:"value"`
	ExpectedCloseDate string     `json:"expected_close_date"`
	PersonID          int        `json:"person_id"`
	OrganizationID    int        `json:"organization_id"`
	IsArchived        bool       `json:"is_archived"`
	ArchiveTime       any        `json:"archive_time"`
	SourceName        string     `json:"source_name"`
	Origin            string     `json:"origin"`
	OriginID          any        `json:"origin_id"`
	Channel           any        `json:"channel"`
	ChannelID         any        `json:"channel_id"`
	WasSeen           bool       `json:"was_seen"`
	NextActivityID    any        `json:"next_activity_id"`
	AddTime           *TimeStamp `json:"add_time"`
	UpdateTime        *TimeStamp `json:"update_time"`
	VisibleTo         string     `json:"visible_to"`
	CcEmail           string     `json:"cc_email"`
}

type LeadAddOptions struct {
	Title             string    `json:"title"`
	OwnerID           int       `json:"owner_id,omitempty"`
	LabelIDs          []string  `json:"label_ids,omitempty"`
	PersonID          int       `json:"person_id,omitempty"`
	OrganizationID    int       `json:"organization_id,omitempty"`
	Value             *Monetary `json:"value,omitempty"`
	ExpectedCloseDate string    `json:"expected_close_data,omitempty"` // YYYY-MM-DD
	VisibileTo        string    `json:"visible_to,omitempty"`
	WasSeen           bool      `json:"was_seen,omitempty"`
	OriginID          string    `json:"origin_id,omitempty"`
	Channel           int       `json:"channel,omitempty"`
	ChannelID         string    `json:"channel_id,omitempty"`
}

type Monetary struct {
	Amount   int    `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}
