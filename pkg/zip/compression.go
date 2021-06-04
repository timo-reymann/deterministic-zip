package zip

import (
	"archive/zip"
	"errors"
)

var ErrInvalidCompressionMethod = errors.New("invalid compression method")

const Deflate = "deflate"
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
