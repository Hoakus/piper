package pipedrive

type Task struct {
	Title            string `json:"title"`
	ID               int    `json:"id"`
	CreatorID        int    `json:"creator_id"`
	Description      string `json:"description"`
	Done             int    `json:"done"`
	DueDate          string `json:"due_date"`
	ParentTaskID     any    `json:"parent_task_id"`
	AssigneeID       int    `json:"assignee_id"`
	AddTime          string `json:"add_time"`
	UpdateTime       string `json:"update_time"`
	MarkedAsDoneTime string `json:"marked_as_done_time"`
	ProjectID        int    `json:"project_id"`
	AdditionalData   struct {
		NextCursor any `json:"next_cursor"`
	} `json:"additional_data"`
}

type TasksGetOptions struct {
	Cursor       string `json:"cursor,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	AssigneeID   int    `json:"assignee_id,omitempty"`
	ProjectID    int    `json:"project_id,omitempty"`
	ParentTaskID int    `json:"parent_task_id,omitempty"`
	Done         int    `json:"done,omitempty"`
}

func (t TasksGetOptions) String() string {
	return Stringify(t)
}

type TaskAddOptions struct {
	ProjectID    int    `json:"project_id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	ParentTaskID int    `json:"parent_task_id,omitempty"` // can not be set to another subtask
	DueDate      string `json:"due_date,omitempty"`       // YYYY-MM-DD
	Done         int    `json:"done,omitempty"`
}

func (t TaskAddOptions) String() string {
	return Stringify(t)
}

type TaskUpdateOptions struct {
	ProjectID    int    `json:"project_id,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	ParentTaskID int    `json:"parent_task_id,omitempty"`
	DueDate      string `json:"due_date,omitempty"` // YYYY-MM-DD
	Done         int    `json:"done,omitempty"`
}

func (t TaskUpdateOptions) String() string {
	return Stringify(t)
}
