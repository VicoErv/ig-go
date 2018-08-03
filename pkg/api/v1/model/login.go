package model

// Login is model of login
type Login struct {
	LoggedInUser User   `json:"logged_in_user"`
	Status       string `json:"status"`
}
