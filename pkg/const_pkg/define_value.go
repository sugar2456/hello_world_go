package const_pkg

type DefineValue int

const (
	UnCategorized DefineValue = iota
	CategoryA
	CategoryB
	CategoryC
)

const (
	const_value1 = 0
	const_value2 = iota + 5
	const_value3
	const_value4
)
