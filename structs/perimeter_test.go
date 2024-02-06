package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {

	rec := Rectangle{
		length: 20.5,
		width:  10.5,
	}

	got := Perimeter(rec)
	want := 62.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rec := Rectangle{
		length: 20.5,
		width:  10.5,
	}

	got := Area(rec)
	want := 215.25

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func ExamplePerimeter() {
	result := Perimeter(Rectangle{5, 10})
	fmt.Println(result)
	// Output: 30
}

func ExampleArea() {
	result := Area(Rectangle{5, 10})
	fmt.Println(result)
	// Output: 50
}

func BenchmarkPerimeter(t *testing.B) {
	rec := Rectangle{5, 10}
	for i := 0; i < t.N; i++ {
		Perimeter(rec)
	}
}

func BenchmarkArea(t *testing.B) {
	rec := Rectangle{5, 10}
	for i := 0; i < t.N; i++ {
		Area(rec)
	}
}
