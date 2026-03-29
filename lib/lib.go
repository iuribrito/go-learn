package lib

import (
	"fmt"

	"github.com/iuribrito/go-learn/lib/internal/foo"
)

var privateVar string = "Private Var!"
var PublicVar string = "Public Var!"

func PrintPrivateVar() {
	fmt.Println(privateVar)
}

func PrintFoo() {
	fmt.Println(foo.Foo)
}
