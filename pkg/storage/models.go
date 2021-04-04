package storage

// User storage model.
type UserInfo struct {
	ID       int
	Email    string
	Password string
	IsActive bool
}

// User storage filter.
type UserFilter struct {
	ID       *int
	Email    *string
	Password *string
}

type TaskInfo struct {
	ID       int
	UserID   int
	Title    string
	Describe string
	Deadline int
}

type TaskFilter struct {
	ID     *int
	UserID int
	From   *int
	To     *int
}
