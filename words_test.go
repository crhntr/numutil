package numutil_test

import (
	"math"
	"testing"

	"github.com/crhntr/numutil"
)

func TestToWords(t *testing.T) {
	// t.SkipNow()
	for n := 0; n < math.MaxInt64; n += ((n + 1) * -3) {
		numutil.ToWords(n)
	}
}
