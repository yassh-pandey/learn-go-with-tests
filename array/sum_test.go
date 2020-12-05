package array

import "testing"

func TestSum(t *testing.T) {
	got := Sum([]float64{17, 99, 133})
	expect := 249.0

	if got != expect {
		t.Errorf("Expected sum: %.2f but got: %.2f", expect, got)
	}
}
