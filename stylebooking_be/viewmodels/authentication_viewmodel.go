package viewmodels

type Token struct {
	Authentication string `json:"authentication,omitempty"`
	Refresh        string `json:"refresh,omitempty"`
}
