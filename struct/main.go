package main

import (
	"fmt"
)

type User struct {
	Name     string
	Email    string
	Password string
}

var users []User

func Register(user User) {
	users = append(users, user)
}

func Get(users []User) {
	for _, data := range users {
		fmt.Printf("Name: %s, Email: %s, Password: %s\n", data.Name, data.Email, data.Password)
	}

	fmt.Println("--------------------------------------------")
}

func main() {

	user1 := User{Name: "John Doe", Email: "johndoe@example.com", Password: "secretpassword"}
	user2 := User{Name: "ROnald Doe", Email: "johndoe@example.com", Password: "secretpassword"}
	// Register the users
	Register(user1)

	Get(users)

	Register(user2)

	Get(users)
}
