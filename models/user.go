package models

type (
	User struct {
		ID        int      `json:"id"`
		Email     string   `json:"email"`
		Password  string   `json:"password,omitempty"`
		Token     string   `json:"token,omitempty"`
		Followers []string `json:"followers,omitempty" `
	}
)
