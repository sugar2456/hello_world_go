package counter

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	c := Counter{}
	if c.total != 0 {
		t.Errorf("Expected initial total to be 0, got %d", c.total)
	}
	c.Increment()
	if c.total != 1 {
		t.Errorf("Expected total to be 1 after increment, got %d", c.total)
	}
	if c.lastUpdate.IsZero() {
		t.Error("Expected lastUpdate to be set after increment")
	}
}

func TestDoUpdateWrong(t *testing.T) {
	c := Counter{}
	initialTotal := c.total
	initialTime := c.lastUpdate

	doUpdateWrong(c)

	if c.total != initialTotal {
		t.Errorf("Expected total to remain %d, got %d", initialTotal, c.total)
	}
	if !c.lastUpdate.Equal(initialTime) {
		t.Errorf("Expected lastUpdate to remain %v, got %v", initialTime, c.lastUpdate)
	}
}
