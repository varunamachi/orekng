package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/varunamachi/orekng/data"
	_ "github.com/varunamachi/orekng/data/sqlite"
)

func main() {
	fmt.Println("Hello Orek!")
}
