package postgresql

// User storage model.
type UserInfo struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	IsActive bool   `db:"is_active"`
}
