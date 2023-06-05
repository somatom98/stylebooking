package viewmodels

type SignInRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignInResponse struct {
	Token string `json:"token,omitempty"`
}
