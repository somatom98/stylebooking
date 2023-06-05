package viewmodels

type SignUpRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type SignUpResponse struct {
	ID string `json:"id,omitempty"`
}
