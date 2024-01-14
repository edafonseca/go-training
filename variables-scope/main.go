package main

import (
	"fmt"

	"github.com/edafonseca/go-training/variables-scope/scope"
)

var name = "Alice"

func main() {
	fmt.Println(name)
	fmt.Println(scope.ScopeVar)
}
