package challenge

type RuleSet struct {
}

// NewRuleSet ...
func NewRuleSet() *RuleSet {
	return &RuleSet{}
}

// AddDep ...
func (rs *RuleSet) AddDep(a, b string) {

}

// AddConflict ...
func (rs *RuleSet) AddConflict(a, b string) {

}

// IsCoherent ...
func (rs *RuleSet) IsCoherent() bool {
	return true
}
