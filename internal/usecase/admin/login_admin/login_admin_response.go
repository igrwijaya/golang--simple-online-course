package login_admin

import "time"

type LoginAdminResponse struct {
	AccessToken  string    `json:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Type         string    `json:"type,omitempty"`
	ExpiredAt    time.Time `json:"expired_at"`
	Scope        string    `json:"scope,omitempty"`
}
