package features

import (
	"errors"
	"testing"
)

func mockErr(msg string) *error {
	err := errors.New(msg)
	return &err
}

func TestFeatures(t *testing.T) {
	if Features() == nil {
		t.Fatal("Features should NEVER return nil")
	}
}
