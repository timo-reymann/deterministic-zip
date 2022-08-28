package zip

import (
	"archive/zip"
	"errors"
)

// ErrInvalidCompressionMethod occurs when an invalid compression method is chosen
var ErrInvalidCompressionMethod = errors.New("invalid compression method")

// Deflate algorithm
const Deflate = "deflate"

// Store algorithm
const Store = "store"

// GetCompressionMethod validates the spec and returns the zip compressor
func GetCompressionMethod(spec string) (uint16, error) {
	switch spec {
	case Deflate:
		return zip.Deflate, nil

	case Store:
		return zip.Store, nil

	default:
		return 0, ErrInvalidCompressionMethod
	}
}
