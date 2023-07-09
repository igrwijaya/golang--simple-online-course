package update_admin

type UpdateAdminRequest struct {
	Name  string `json:"name,omitempty" binding:"required" json:"name,omitempty"`
	Email string `json:"email,omitempty" binding:"required,email" json:"email,omitempty"`
}
