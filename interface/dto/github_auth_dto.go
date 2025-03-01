package dto

type GitHubAuthRequest struct {
	Token string `json:"token"`
}
type GitHubAuthResponse struct {
	UserID int `json:"userId"`
}

type GitHubUser struct {
	Login string `json:"login"` // loginにユーザ名が入る
}

type GitHubEmail struct {
	Email string `json:"email"`
}