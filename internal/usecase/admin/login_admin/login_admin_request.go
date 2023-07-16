package login_admin

type LoginAdminRequest struct {
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}
