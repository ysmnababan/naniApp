package dto

type UserEditReq struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Picture_URL string `json:"picture_url"`
}
