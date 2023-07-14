package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ImageUrl string `json:"image_url"`
}

var users map[int]User

func NewUserDB() {
	users = make(map[int]User)
}

func CreateUser(u *User) {
	id := len(users) + 1
	u.ID = id
	users[id] = *u
}

func GetUser(id int) User {
	return users[id]
}

func UpdateUser(id int, update map[string]string) User {
	user := users[id]

	for key, value := range update {
		switch key {
		case "email":
			user.Email = value
		case "password":
			user.Password = value
		case "image_url":
			user.ImageUrl = value
		}
	}

	users[id] = user
	return users[id]
}
