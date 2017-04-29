package main

import (
	"errors"
	"fmt"

	"github.com/gopherci/importsC/testdata/indirect/direct"
)

func main() {
	print("does not import C\n")
	fmt.Println("does not import C\n")

	fmt.Println(errors.New("sdf"))
	direct.Foo()
}
