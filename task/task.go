package task
// This package contains all data and methods related to
// RapptiveTask tasks. This package is responsible for
// all CRUD actions and other functionalities ralted to
// RapptiveTask tasks.

// Bind the request JSON object to RapptiveTask struct
// This RapptiveTask struct currently holds these values:
// TaskSubject: A subject of the task (gets displayed in the frontend)
// TaskContent: The content of the task (")
// TaskID: A unique identifier for a task
type RapptiveTask struct {
	TaskSubject string `json:"subject" binding:"required"`
	TaskContent string `json:"content" binding:"required"`
	TaskID int `json:"id" binding:"required"`
}

func (t *RapptiveTask) Create() bool {
    return false
}

func (t *RapptiveTask) Read() bool {
    return false
}

func (t *RapptiveTask) Update() bool {
    return false
}

func (t *RapptiveTask) Delete() bool {
    return false
}
