package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, s := range areaTests {
		got := s.shape.Area()
		if got != s.want {
			t.Errorf("got %g want %g", got, s.want)
		}
	}
}
