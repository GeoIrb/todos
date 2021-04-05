package postgresql

type userInfo struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	IsActive bool   `db:"is_active"`
}

type taskInfo struct {
	ID          int    `db:"id"`
	UserID      int    `db:"user_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Deadline    int    `db:"deadline"`
}
