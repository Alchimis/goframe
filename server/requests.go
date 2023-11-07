package server

type LoginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
