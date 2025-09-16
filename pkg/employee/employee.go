package employee

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("Employee Name: %s, ID: %s", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployee() []Employee {
	newEmployees := []Employee{
		{Name: "Alice", ID: "E001"},
		{Name: "Bob", ID: "E002"},
	}
	return newEmployees
}
