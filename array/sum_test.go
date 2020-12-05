package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum([]float64{17, 99, 133})
	expect := 249.0

	if got != expect {
		t.Errorf("Expected sum: %.2f but got: %.2f", expect, got)
	}
}

func TestSumAll(t *testing.T) {
	arr := [][]float64{{1, 55, 9}, {12, 45, -4}, {-83, 0, 83}}
	got := SumAll(arr)
	expect := []float64{65, 53, 0}
	if len(got) == len(expect) {
		for i, v := range got {
			if expect[i] != v {
				t.Errorf("Expected: %v, got this: %v for input = %v", expect, got, arr)
			}
		}
	} else {
		t.Errorf("Expected: %v, got this: %v for input = %v", expect, got, arr)
	}
}

func TestSumAllTails(t *testing.T) {

	assertCorrectness := func(t *testing.T, got, expect []float64) {
		t.Helper()
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Expected: %v, got this: %v", expect, got)
		}
	}

	t.Run("Tail sum of non empty slices", func(t *testing.T) {
		got := SumAllTails([]float64{12, 44, 9}, []float64{-9, 9, 0})
		expect := []float64{53, 9}
		assertCorrectness(t, got, expect)
	})

	t.Run("Tail sum for empty slices", func(t *testing.T) {
		got := SumAllTails([]float64{}, []float64{-999})
		expect := []float64{0, 0}
		assertCorrectness(t, got, expect)
	})
}
