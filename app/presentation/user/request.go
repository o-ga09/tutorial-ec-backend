package user

type PostUserParam struct {
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	Pref        string `json:"pref,omitempty"`
	City        string `json:"city,omitempty"`
	Extra       string `json:"extra,omitempty"`
}
