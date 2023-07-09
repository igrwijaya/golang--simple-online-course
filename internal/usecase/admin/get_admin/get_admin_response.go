package get_admin

import "golang-online-course/pkg/response"

type AdminDto struct {
	Id    uint   `json:"id,omitempty"`
	Email string `json:"email,omitempty" json:"email,omitempty"`
	Name  string `json:"name,omitempty" json:"name,omitempty"`
}

type GetAdminResponse struct {
	response.Basic
	TotalRecord int64      `json:"total_record,omitempty"`
	CurrentPage int        `json:"current_page,omitempty"`
	Limit       int        `json:"limit,omitempty"`
	Data        []AdminDto `json:"data,omitempty"`
}
