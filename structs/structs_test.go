package structs

import "testing"

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{"Rectangle", Rectangle{Height: 12, Width: 6}, 72.0},
		{"Circle", Circle{Radius: 10}, 314.15927},
		{"Triangle", Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.area
			if got != want {
				t.Errorf("%#v got %g want %g", tt.shape, got, want)
			}
		})
	}
}
