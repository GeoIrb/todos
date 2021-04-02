package http

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token *string `json:"token"`
}

type userInfo struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type getUserListResponse struct {
	Users []userInfo `json:"users"`
}

type registrationRequest struct {
	Email string `json:"email"`
}

type createRequest struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
