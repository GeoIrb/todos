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
