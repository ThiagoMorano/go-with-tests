package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("should return 5 times when called with 5", func(t *testing.T) {
		parameterizedTestRepeat(t, "a", 5, "aaaaa")
	})
	t.Run("should return 4 times when called with 4", func(t *testing.T) {
		parameterizedTestRepeat(t, "a", 4, "aaaa")
	})
}

func parameterizedTestRepeat(t testing.TB, character string, times int, expected string) {
	received := Repeat(character, times)

	if received != expected {
		t.Errorf("expected %q but got %q", expected, received)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
