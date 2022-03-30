package web

type CustomerResponse struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
}
