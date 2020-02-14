package main

import (
	"fmt"

	"github.com/kgolding/go-app-structure/internal/prv"
	"github.com/kgolding/go-app-structure/pkg/one"
)

func main() {
	fmt.Println("This is main")
	one.Hello()
	prv.Hello()
}
