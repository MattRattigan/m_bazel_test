package go_hello_world

import (
	"testing"
)

func TestGreeter(t *testing.T) {
	want := "Hello World!"
	got := HelloWorld()
	if want != got {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}
