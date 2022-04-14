package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	testcases := []struct{ got, want string }{
		{"a", "aaaaa"},
		{"hello", "hellohellohellohellohello"},
		{"meep", "meepmeepmeepmeepmeep"},
		{"", "Invalid input"},
	}

	for _, testcase := range testcases {
		s := Repeat(testcase.got)
		if s != testcase.want {
			t.Errorf("got %q, want %q", s, testcase.want)
		}
	}
}

func ExampleRepeat() {
	ans := Repeat("x")
	fmt.Println(ans)
	// Output: xxxxx
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("meep")
	}
}
