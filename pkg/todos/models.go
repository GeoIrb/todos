package todos

type TaskInfo struct {
	ID       int
	UserID   int
	Title    string
	Describe string
	Deadline int
}

type Filter struct {
	ID     *int
	UserID int
	From   *int
	To     *int
}
