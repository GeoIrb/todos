package storage

// User storage model.
type UserInfo struct {
	ID       string
	Email    string
	Password string
	IsActive bool
}

// User storage filter.
type UserFilter struct {
	ID       *string
	Email    *string
	Password *string
}
