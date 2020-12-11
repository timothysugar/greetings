package greetings

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hi, Sue. Welcome!"
	if got := Hello("Sue"); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
