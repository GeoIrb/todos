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

type Auth struct {
	Token *string
}
