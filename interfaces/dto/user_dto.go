package dto

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
