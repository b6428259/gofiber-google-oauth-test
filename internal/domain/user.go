package domain

type User struct {
    GoogleID string `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
}
