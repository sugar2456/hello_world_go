package counter

import (
	"fmt"
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

func TestDoUpdate(t *testing.T) {
	c := Counter{}
	fmt.Println("main1:", c.String())
	doUpdateWrong(c)
	fmt.Println("main2:", c.String())
	doUpdateRight(&c)
	fmt.Println("main3:", c.String())
	doUpdateRight(&c)
	fmt.Println("main4:", c.String())
	if c.total != 2 {
		t.Errorf("Expected total to be 2 after two increments, got %d", c.total)
	}
}
