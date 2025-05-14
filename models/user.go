package models

// UserProfile represents user profile data
type UserProfile struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Emoji    string `json:"emoji"`
}
