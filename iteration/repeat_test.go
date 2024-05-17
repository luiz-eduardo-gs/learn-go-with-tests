package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 3)
	expected := "aaa"

	if expected != actual {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
