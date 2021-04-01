package user

// Registartion info.
type Registartion struct {
	Email string
}

// Login info.
type Login struct {
	Email    string
	Password string
}

// Auth info.
type Auth struct {
	Token *string
}

// Create user info.
type Create struct {
	Email       string
	OldPassword string
	NewPassword string
}

// Filter user.
type Filter struct {
	ID       *string
	Email    *string
}

type User struct {
	ID    string
	Email string
}
