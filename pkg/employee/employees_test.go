package employee

import (
	"fmt"
	"testing"
)

func TestManger(t *testing.T) {
	m := Manager{
		Employee: Employee{Name: "John", ID: "M001"},
		Reports:  []Employee{},
	}

	fmt.Println(m.ID)
	fmt.Println(m.Name)
	fmt.Println(m.Description())

	fmt.Println(m.Employee)
	fmt.Println(m.Reports)

	m.Reports = m.FindNewEmployee()
	fmt.Println(m.Employee)
	fmt.Println(m.Reports)
}
