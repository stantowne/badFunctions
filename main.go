package main

import (
	"fmt"
	"github.com/stantowne/badFunctions/badAdder"
	"github.com/stantowne/badFunctions/badOther"
)

func main() {
	a := 1
	b := 5
	c := badAdder.BadAdder(a, b)
	d := badAdder.VeryBadAdder(a, b)
	e := badOther.BadMultiplier(a, b)
	fmt.Println(a, b, c, d, e)

}
