package create_admin

type CreateAdminRequest struct {
	Name            string `json:"name,omitempty" binding:"required"`
	Email           string `json:"email,omitempty" binding:"required,email"`
	Password        string `json:"password,omitempty" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password,omitempty" binding:"required,min=8"`
}
