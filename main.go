package main

import (
	"fmt"

	"github.com/lutffmn/clay/clay"
)

func dummyHandler1() {
	fmt.Println("dummy 1")
}

func dummyHandler2() {
	fmt.Println("dummy 2")
}

func main() {
	app := clay.New("dummy app", "this is a dummy app", "dummy [OPTION]", "v0")
	app.Short("d", dummyHandler1, false)
	app.Long("d", dummyHandler2, false)

	fmt.Println(app)
}
