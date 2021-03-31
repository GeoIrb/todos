package database

// User storage model.
type UserInfo struct {
	ID       string
	Username string
	Password string
}

// User database filter.
type UserFilter struct {
	ID string
}
