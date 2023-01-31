package authentication

type Password struct {
	Email       string `json:"Email"`
	OldPassword string `json:"OldPassword"`
	NewPassword string `json:"NewPassword"`
}
