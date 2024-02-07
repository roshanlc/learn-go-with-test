package interfaces

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Rectangle", func(t *testing.T) {
		rec := Rectangle{Length: 20.0, Width: 10.0}

		got := rec.Area()
		want := 200.0

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		cir := Circle{Radius: 10.0}

		got := cir.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

}

func ExampleArea() {
	rec := Rectangle{Length: 20.0, Width: 10.0}
	area := rec.Area()
	fmt.Println(area)

	cir := Circle{Radius: 10.0}
	area = cir.Area()
	fmt.Println(area)
	// Output:
	// 200
	// 314.1592653589793
}

func BenchmarkArea(b *testing.B) {
	b.Run("Benchmark for rectangle Area", func(b *testing.B) {
		rec := Rectangle{Length: 20.0, Width: 10.0}

		for i := 0; i < b.N; i++ {
			rec.Area()
		}
	})

	b.Run("Benchmark for circle Area", func(b *testing.B) {
		cir := Circle{Radius: 10.0}

		for i := 0; i < b.N; i++ {
			cir.Area()
		}
	})

}
