package database

// User storage model.
type UserInfo struct {
	ID       string
	Email    string
	Username string
	Password string
}

// User database filter.
type UserFilter struct {
	ID string
}
