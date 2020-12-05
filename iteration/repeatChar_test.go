package iteration

import (
	"fmt"
	"testing"
)

func TestRepeatChars(t *testing.T) {
	got := RepeatChars("a", 5)
	expect := "aaaaa"
	if got != expect {
		t.Errorf("Ecpected: %q Got: %q", expect, got)
	}

}
func BenchmarkRepeatChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatChars("a", 5)
	}
}

func ExampleRepeatChars() {
	repeatedChars := RepeatChars("d", 8)
	fmt.Println(repeatedChars)
	// Output: dddddddd
}
