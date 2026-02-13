# Hello Modules

## Instalation

Execute the following command:
```bash
go get -u github.com/sarreolam/hello_modules
```

## Usage

```go

package main

import (
	"fmt"

	"github.com/sarreolam/hello_modules"
)

func main() {
	var message string
	message = hello_modules.Hello1("Santiago")
	fmt.Println(message)

	fmt.Printf(hello_modules.RandomHello(), "Santiago")
}
```