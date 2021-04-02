package postgresql

type userInfo struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	IsActive bool   `db:"is_active"`
}
