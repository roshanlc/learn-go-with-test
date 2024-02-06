package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("A")
	want := "AAAAA"

	if got != want {
		t.Errorf("expected %q, got %q", got, want)
	}
}

func ExampleRepeat() {
	repeated := Repeat("R")
	fmt.Println(repeated)
	// Output: RRRRR
}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("A")
	}
}
