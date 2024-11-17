package mod

type LoginResponse struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
