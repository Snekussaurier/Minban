package mod

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type IdRequest struct {
	ID int `json:"id" binding:"required"`
}

type NameRequest struct {
	Name string `json:"name" binding:"required"`
}
