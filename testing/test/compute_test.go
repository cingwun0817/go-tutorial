package compute

import (
	compute "test/pkg"
	"testing"
)

func TestPlusElements(t *testing.T) {
	if compute.Plus(1, 1) != 2 {
		t.Error("1 + 1 did not equal 2")
	}
}

func TestTimesElements(t *testing.T) {
	if compute.Times(2, 2) != 4 {
		t.Error("2 * 2 did not equal 4")
	}
}
