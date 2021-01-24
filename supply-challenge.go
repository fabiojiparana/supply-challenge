package main

import (
	"fmt"

	"github.com/supply-challenge/rule"
)

func main() {
	rs := rule.NewRuleSet()
	rs.AddDep("a", "b")
	rs.AddDep("b", "c")
	rs.AddDep("c", "d")
	rs.AddDep("d", "e")
	rs.AddDep("a", "f")
	rs.AddConflict("e", "f")

	if rs.IsCoherent() {
		fmt.Println("s.IsCoherent failed")
	} else {
		fmt.Println("Coherent")
	}
}
