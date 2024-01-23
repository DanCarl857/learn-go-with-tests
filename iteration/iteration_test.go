package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	repeated := Iteration("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a", 5)
	}
}

func ExampleIteration() {

	ans := Iteration("a", 5)
	fmt.Println(ans)
	// Output: aaaaa
}
