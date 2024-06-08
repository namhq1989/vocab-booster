package dto

type GetAccessTokenByStaffIDRequest struct {
	StaffID string `query:"staffId" validate:"required" message:"staff_invalid_id"`
}

type GetAccessTokenByStaffIDResponse struct {
	AccessToken string `json:"accessToken"`
}
