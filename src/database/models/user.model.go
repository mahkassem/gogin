package models

type User struct {
	Entity

	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	// Metadata string
}
