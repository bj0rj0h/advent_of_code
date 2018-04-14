package three

type policy interface {
	getNextValue() policy
	getValue() int
}

type incrementalPolicy struct {
	value int
}

func (p incrementalPolicy) getNextValue() policy {
	result := incrementalPolicy{p.value + 1}
	return result
}

func (p incrementalPolicy) getValue() int {
	return p.value
}
