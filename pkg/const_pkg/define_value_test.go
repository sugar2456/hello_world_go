package const_pkg

import (
	"fmt"
	"testing"
)

func TestDefineValue(t *testing.T) {
	fmt.Println(UnCategorized, CategoryA, CategoryB, CategoryC)
}

func TestConstValue(t *testing.T) {
	fmt.Println(const_value1, const_value2, const_value3, const_value4)
}
