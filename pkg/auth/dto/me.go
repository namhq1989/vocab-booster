package dto

type GetMeRequest struct{}

type GetMeResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
