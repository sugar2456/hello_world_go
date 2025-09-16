package counter

import (
	"fmt"
	"time"
)

type Counter struct {
	total      int
	lastUpdate time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdate = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, lastUpdate: %v", c.total, c.lastUpdate)
}

func doUpdateWrong(c Counter) {
	c.Increment()
}
