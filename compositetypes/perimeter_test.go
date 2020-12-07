package compositetypes

import "testing"

func TestPerimeter(t *testing.T) {
	assert := func(t *testing.T, expect, got float64) {
		t.Helper()
		if expect != got {
			t.Errorf("Expected %g but got: %g", expect, got)
		}
	}
	t.Run("Rect perimeter", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  30.0,
			Height: 50.0,
		}
		got := rectangle.Perimeter()
		expect := 160.0
		assert(t, expect, got)
	})
	t.Run("Circle perimeter", func(t *testing.T) {
		circle := Circle{
			Radius: 15.0,
		}
		got := circle.Perimeter()
		expect := 94.24777960769379
		assert(t, expect, got)
	})
}

/*
	If you see the TestArea test below then it's has a lot of code repeatition.
	We want to solve this problem by using Go Interfaces.
*/

// func TestArea(t *testing.T) {
// 	assert := func(t *testing.T, expect, got float64) {
// 		t.Helper()
// 		if expect != got {
// 			t.Errorf("Expected %g but got: %g", expect, got)
// 		}
// 	}
// 	t.Run("Rectangle area", func(t *testing.T) {
// 		rectangle := Rectangle{
// 			Width:  30.0,
// 			Height: 15.0,
// 		}
// 		got := rectangle.Area()
// 		expect := 450.0
// 		assert(t, expect, got)
// 	})
// 	t.Run("Circle area", func(t *testing.T) {
// 		circle := Circle{
// 			Radius: 10.0,
// 		}
// 		got := circle.Area()
// 		expect := 314.1592653589793
// 		assert(t, expect, got)
// 	})
// }

// Notice how concise has the code testing code become by making use of interfaces 
func TestArea(t *testing.T) {
	testArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Expected: %g but got: %g", got, want)
		}
	}
	t.Run("Rectangle area test", func(t *testing.T) {
		rect := Rectangle{30, 17}
		testArea(t, rect, 510.0)
	})
	t.Run("Circle area test", func(t *testing.T) {
		circle := Circle{15}
		testArea(t, circle, 706.8583470577034)
	})
}
