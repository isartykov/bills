package user

type User struct {
	id       int    `json:"id"`
	name     string `json:"name"`
	password string `json:"password"`
}
