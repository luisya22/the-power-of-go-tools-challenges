package count_test

import (
	"bytes"
	count "github.com/luisya22/the-power-of-go-tools-challenges"
	"testing"
	"time"
)

func TestReturnNextNumber(t *testing.T) {
	t.Parallel()

	c := count.Counter{Value: 5}
	num := c.Next()

	want := 6

	if want != num {
		t.Errorf("want: %v, got: %v", want, num)
	}
}

func TestPrintNumbersToTerminal(t *testing.T) {
	t.Parallel()

	fakeTerminal := &bytes.Buffer{}
	c := &count.Counter{
		Printer:  fakeTerminal,
		Limit:    5,
		WaitTime: 1 * time.Nanosecond,
	}

	c.Run()

	want := "1\n2\n3\n4\n5\n"
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("want: %q, got: %q", want, got)
	}
}
