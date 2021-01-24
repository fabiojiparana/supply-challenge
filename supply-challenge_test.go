package main

import (
	"testing"

	"github.com/supply-challenge/rule"
)

const (
	errorNotCoherent = "failed > is not coherent"
)

func TestDependsAA(t *testing.T) {
	rs := rule.NewRuleSet()
	rs.AddDep("a", "a")

	if !rs.IsCoherent() {
		t.Error(errorNotCoherent)
	}
}

func TestDependsAB_BA(t *testing.T) {
	rs := rule.NewRuleSet()
	rs.AddDep("a", "b")
	rs.AddDep("b", "a")

	if !rs.IsCoherent() {
		t.Error(errorNotCoherent)
	}
}

func TestExclusiveAB(t *testing.T) {
	rs := rule.NewRuleSet()
	rs.AddDep("a", "b")
	rs.AddConflict("a", "b")

	if rs.IsCoherent() {
		t.Error(errorNotCoherent)
	}
}

func TestExclusiveAB_BC(t *testing.T) {
	rs := rule.NewRuleSet()
	rs.AddDep("a", "b")
	rs.AddDep("b", "c")
	rs.AddConflict("a", "c")

	if rs.IsCoherent() {
		t.Error(errorNotCoherent)
	}
}
func TestDeepDeps(t *testing.T) {
	rs := rule.NewRuleSet()
	rs.AddDep("a", "b")
	rs.AddDep("b", "c")
	rs.AddDep("c", "d")
	rs.AddDep("d", "e")
	rs.AddDep("a", "f")
	rs.AddConflict("e", "f")

	if rs.IsCoherent() {
		t.Error(errorNotCoherent)
	}
}
