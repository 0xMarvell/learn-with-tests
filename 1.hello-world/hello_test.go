package hello

import "testing"

func TestHello(t *testing.T) {
	testcases := []struct{ got, want string }{
		{"Mary", "Hello, Mary"},
		{"Marvellous", "Hello, Marvellous"},
		{"bell", "Hello, bell"},
		{"", "Hello"},
	}
	for _, testcase := range testcases {
		s := Hello(testcase.got)
		if s != testcase.want {
			t.Fatalf("Hello(%q) = %q, want %q", testcase.got, s, testcase.want)
		}
	}
}
