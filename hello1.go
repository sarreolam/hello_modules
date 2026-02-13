package hello_modules

import (
	"fmt"
	"math/rand"
)

func Hello1(name string) string {
	message := fmt.Sprintf("Hello 1, %s from hello1", name)
	return message
}

func RandomHello() string {
	greetings := []string{
		"Hello %v",
		"Hi %v",
		"Hey %v",
		"Greetings %v",
		"Salutations %v",
	}
	return greetings[rand.Intn(len(greetings))]
}
