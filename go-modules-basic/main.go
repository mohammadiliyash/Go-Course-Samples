package main

import (
	"fmt"
	m "go-modules-basic/account/models"
)

func main() {

	i := 10

	u := m.User{}
	u.Email = "hi"

	fmt.Println("Hi I am here", i, u)
}
