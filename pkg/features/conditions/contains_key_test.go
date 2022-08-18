package conditions

import (
	"testing"
)

func TestContainsKey(t *testing.T) {
	val := map[string]string{
		"zzzzz":   "zzzzz",
		"element": "element",
	}

	if !ContainsKey(&val, "element") {
		t.Fatal("Contains check not working")
	}
	if ContainsKey(&val, "missing element") {
		t.Fatal("Contains check found element not in the slice")
	}
}
