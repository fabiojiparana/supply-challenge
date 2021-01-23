package rules

type RuleSet struct {
	rulers    map[string][]string
	conflicts map[string][]string
}

// NewRuleSet create a new instance
func NewRuleSet() *RuleSet {
	return &RuleSet{
		rulers:    make(map[string][]string),
		conflicts: make(map[string][]string),
	}
}

// AddDep adds new dependency between A and B
func (rs *RuleSet) AddDep(a, b string) {
	rs.rulers[a] = append(rs.rulers[a], b)
	rs.rulers[b] = append(rs.rulers[b], a)
}

// AddConflict adds new conflict between A and B
func (rs *RuleSet) AddConflict(a, b string) {
	rs.conflicts[a] = append(rs.conflicts[a], b)
	rs.conflicts[b] = append(rs.conflicts[b], a)
}

// IsCoherent checks the set of coherent rules
func (rs *RuleSet) IsCoherent() bool {
	for a, values := range rs.conflicts {
		for _, b := range values {
			if rs.checkConnectionBetweenValues(a, b) {
				return false
			}
		}
	}

	return true
}

// checkConnectionBetweenValues checks if there is a connection between the values A and B
func (rs *RuleSet) checkConnectionBetweenValues(a, b string) bool {
	checked := make(map[string]bool)
	var values []string
	checked[a] = true

	values = append(values, rs.rulers[a]...)

	for len(values) > 0 {
		firstNode := values[0]
		values = values[1:]

		if firstNode == b {
			return true
		}

		if !checked[firstNode] {
			values = append(values, rs.rulers[firstNode]...)
			checked[firstNode] = true
		}
	}

	return false
}
