package integers

import (
	"fmt"
	"testing"
)

func ExampleAdder() {
	sum := Adder(10, 20)
	fmt.Println(sum)
	// Output: 30
}
func TestAdder(t *testing.T) {
	sum := Adder(2, 2)
	expect := 4
	if sum != expect {
		t.Errorf("Expected %d but got %d", expect, sum)
	}
}
