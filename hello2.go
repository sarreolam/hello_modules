package hello_modules

import (
	"fmt"
)

func Hello2(name string) string {
	message := fmt.Sprintf("Hello 1, %s from hello2", name)
	return message
}
