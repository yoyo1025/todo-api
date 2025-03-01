package dto

type GitHubAuthRequest struct {
	Token string `json:"token"`
}
type GitHubAuthResponse struct {
	UserID int `json:"userId"`
}