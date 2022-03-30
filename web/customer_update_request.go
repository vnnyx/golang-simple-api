package web

type CustomerUpdateRequest struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `validate:"max=200" json:"username,omitempty"`
	Password string `validate:"max=200" json:"password,omitempty"`
}
