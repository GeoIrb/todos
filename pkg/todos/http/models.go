package http

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    int    `json:"deadline"`
}

type updateTaskRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    int    `json:"deadline"`
}

type deleteTaskRequest struct {
	ID int `json:"id"`
}

type taskInfo struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    int    `json:"deadline"`
}

type getTaskListResponse struct {
	Tasks []taskInfo `json:"task"`
}
