package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ImageUrl string `json:"image_url"`
}

var users map[string]User

func (u *User) CreateUser() {
	users = make(map[string]User) // allocate memory to the map

	users[u.Email] = *u
}

func (u *User) GetUser(email string) User {
	return users[email]
}

func (u *User) UpdateUser() {
	users[u.Email] = *u
}
