package tree

import "testing"

func TestIntTree(t *testing.T) {
	var tree *IntTree
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		tree = tree.Insert(v)
	}

	tests := []struct {
		value    int
		expected bool
	}{
		{5, true},
		{3, true},
		{7, true},
		{2, true},
		{4, true},
		{6, true},
		{8, true},
		{1, false},
		{9, false},
	}

	for _, test := range tests {
		if result := tree.Contains(test.value); result != test.expected {
			t.Errorf("Contains(%d) = %v; want %v", test.value, result, test.expected)
		}
	}
}
