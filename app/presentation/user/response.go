package user

type getUserResponse struct {
	User userResponseModel
}

type userResponseModel struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	Address     string `json:"address,omitempty"`
}
