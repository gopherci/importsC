package direct

import "C"
import (
	"errors"
	"fmt"
)

func Foo() {
	fmt.Println("does not import C\n")
	print("imports C\n")
	fmt.Println(errors.New("sdf"))
}
