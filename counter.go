package count

import (
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

type Counter struct {
	Value    int
	Printer  io.Writer
	Limit    int
	WaitTime time.Duration
}

func (c *Counter) Next() int {
	c.Value += 1
	return c.Value
}

func (c *Counter) Run() {
	for c.Limit > c.Value {
		time.Sleep(c.WaitTime)
		fmt.Fprintf(c.Printer, "%v\n", c.Next())
	}
}

func NewCounter() *Counter {
	return &Counter{
		Value:    -1,
		Printer:  os.Stdout,
		Limit:    math.MaxInt,
		WaitTime: 10 * time.Minute,
	}
}
