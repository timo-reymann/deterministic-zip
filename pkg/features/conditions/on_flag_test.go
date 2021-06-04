package conditions

import "testing"

func TestOnFlag(t *testing.T) {
	if !OnFlag(true) {
		t.Fatal("True should lead to onFlag")
	}

	if OnFlag(false) {
		t.Fatal("False should not lead to onFlag")
	}
}
