package pipedrive

type DealFields struct {
	ID                  int        `json:"id"`
	Key                 string     `json:"key"`
	Name                string     `json:"name"`
	OrderNumber         int        `json:"order_nr"`
	FieldType           string     `json:"field_type"`
	AddTime             *TimeStamp `json:"add_time"`
	UpdateTime          *TimeStamp `json:"update_time"`
	LastUpdatedByUserID int        `json:"last_updated_by_user_id"`
	ActiveFlag          bool       `json:"active_flag"`
	EditFlag            bool       `json:"edit_flag"`
	IndexVisibleFlag    bool       `json:"index_visible_flag"`
	DetailsVisibleFlag  bool       `json:"details_visible_flag"`
	AddVisibleFlag      bool       `json:"add_visible_flag"`
	ImportantFlag       bool       `json:"important_flag"`
	BulkEditAllowed     bool       `json:"bulk_edit_allowed"`
	SearchableFlag      bool       `json:"searchable_flag"`
	FilteringAllowed    bool       `json:"filtering_allowed"`
	SortableFlag        bool       `json:"sortable_flag"`
	Options             any        `json:"options"`
	MandatoryFlag       any        `json:"mandatory_flag"`
}

type GetDealsFieldsOptions struct {
	Start int `url:"start,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func (p GetDealsFieldsOptions) String() string {
	return Stringify(p)
}
