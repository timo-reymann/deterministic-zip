package conditions

import "testing"

func TestHasElements(t *testing.T) {
	val := []string{
		"",
	}
	if !HasElements(&val) {
		t.Fatal("Empty check not working")
	}
}
