package main

import (
	"fmt"

	_ "github.com/varunamachi/orekng/data"
	_ "github.com/varunamachi/orekng/data/sqlite"
	_ "github.com/varunamachi/orekng/rest"
)

func main() {
	fmt.Println("Hello Orek!")
}
