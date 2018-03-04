package main

import (
	"fmt"

	"github.com/yanniqq/stringutil"
)

func main() {
	fmt.Printf("Hello, GO world!\n")
	fmt.Printf("And now in reverse order:")
	fmt.Printf(stringutil.Reverse("Hello, GO world!\n"))
}