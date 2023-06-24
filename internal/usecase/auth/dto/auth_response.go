package dto

type AuthResponseDto struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Type         string `json:"type,omitempty"`
	ExpiredAt    string `json:"expired_at,omitempty"`
	Scope        string `json:"scope,omitempty"`
}
