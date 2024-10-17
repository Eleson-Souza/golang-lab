package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	user := User{Name: "Eleson", Email: "eleson@domain.com", Age: 27}

	userJSON, _ := json.Marshal(user)

	fmt.Println(string(userJSON))
}
