package hello_test

import (
	"bytes"
	"testing"
)

func TestPrintsHelloMessageToTerminal(t *testing.T) {
	t.Parallel()

	fakeTerminal := &bytes.Buffer{}
	hello.PrintTo(fakeTerminal)

	want := "Hello, world"
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("want: %q, got: %q")
	}
}
