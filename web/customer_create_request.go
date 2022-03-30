package web

type CustomerCreateRequest struct {
	Name     string `validate:"required" json:"name,omitempty"`
	Username string `validate:"required,min=1,max=200" json:"username,omitempty"`
	Password string `validate:"required,min=8,max=200" json:"password,omitempty"`
}
