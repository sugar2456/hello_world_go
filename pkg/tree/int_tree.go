package tree

type IntTree struct {
	Left  *IntTree
	Value int
	Right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{Value: val}
	}
	if val < it.Value {
		it.Left = it.Left.Insert(val)
	} else {
		it.Right = it.Right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	if it == nil {
		return false
	}
	if val == it.Value {
		return true
	} else if val < it.Value {
		return it.Left.Contains(val)
	} else {
		return it.Right.Contains(val)
	}
}
