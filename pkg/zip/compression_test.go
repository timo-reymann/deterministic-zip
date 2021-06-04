package zip

import (
	"archive/zip"
	"testing"
)

func TestGetCompressionMethod(t *testing.T) {
	testCases := []struct {
		method         string
		expectedMethod uint16
	}{
		{
			method:         Store,
			expectedMethod: zip.Store,
		},
		{
			method:         Deflate,
			expectedMethod: zip.Deflate,
		},
	}

	for _, tc := range testCases {
		m, _ := GetCompressionMethod(tc.method)
		if m != tc.expectedMethod {
			t.Fatalf("Expected method %d, but got %d for input %s", tc.expectedMethod, m, tc.method)
		}
	}
}
