package numutil_test

import (
	"testing"

	"github.com/crhntr/numutil"
)

func TestToWords(t *testing.T) {
	for n := -19; n < 100; n++ {
		t.Log(numutil.ToWords(n))
	}

	for n := 100; n < 1000; n += 55 {
		t.Log(numutil.ToWords(n))
	}

	for n := 1000; n < 1000000; n += 77654 {
		t.Log(numutil.ToWords(n))
	}

	for n := 1000000; n < 1000000000; n += 7747654 {
		t.Log(numutil.ToWords(n))
	}

	for n := 1000000000; n < 1000000000000; n += 7747543654 {
		t.Log(numutil.ToWords(n))
	}
}
