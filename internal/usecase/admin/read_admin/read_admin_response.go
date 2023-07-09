package read_admin

import "golang-online-course/pkg/response"

type ReadAdminResponse struct {
	response.Basic
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}
