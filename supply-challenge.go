package main

import (
	"fmt"

	"github.com/supply-challenge/rules"
)

func main() {
	rs := rules.NewRuleSet()
	rs.AddDep("a", "b")
	rs.AddDep("b", "c")
	rs.AddConflict("a", "c")

	if rs.IsCoherent() {
		fmt.Println("s.IsCoherent failed")
	} else {
		fmt.Println("Coherent")
	}
}
