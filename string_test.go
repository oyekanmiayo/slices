package slices

import "testing"

func TestString_Contains(t *testing.T) {
	str := String{}
	if got := str.Contains([]string{"a", "b"}, "a"); got != true {
		t.Errorf("Contains() = %v, want %v", got, true)
	}
}
