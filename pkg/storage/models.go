package storage

// User storage model.
type UserInfo struct {
	ID       string
	Email    string
	Username string
	Password string
}

// User storage filter.
type UserFilter struct {
	ID string
}
