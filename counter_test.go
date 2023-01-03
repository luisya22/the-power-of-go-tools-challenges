package counter_test

import (
	"testing"
	counter "the-power-of-go-tools-challenges"
)

func TestPrintsNextNumber(t *testing.T) {
	t.Parallel()

	counter.Next()
}
