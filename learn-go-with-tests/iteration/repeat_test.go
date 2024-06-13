package iteration

import "testing"

const repeateCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeateCount; i++ {
		repeated += character
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}