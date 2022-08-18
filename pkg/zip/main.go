package zip

import (
	"archive/zip"
	"compress/flate"
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

const extension = ".zip"

var extra = make([]byte, 0)

// ModifiedTimestamp contains the default modification timestamp used for all files in the archive
var ModifiedTimestamp = time.Date(2018, 11, 01, 0, 0, 0, 0, time.UTC)

func createFileName(input string) string {
	if strings.HasSuffix(input, extension) {
		return input
	}
	return input + extension
}

// Create new zip file with the given configuration and compression method
func Create(c *cli.Configuration, compression uint16) error {
	finalName := createFileName(c.ZipFile)

	sort.Strings(c.SourceFiles)

	newZipFile, err := os.OpenFile(finalName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	c.ZipFile = finalName

	zipWriter := zip.NewWriter(newZipFile)
	registerCompressors(zipWriter)

	for _, srcFile := range c.SourceFiles {
		if err := appendFile(srcFile, zipWriter, compression); err != nil {
			return err
		}
	}

	return zipWriter.Close()
}

func registerCompressors(zipWriter *zip.Writer) {
	zipWriter.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w, flate.BestSpeed)
	})
}

func appendFile(srcFile string, zipWriter *zip.Writer, compression uint16) error {
	output.Infof("Adding file %s", srcFile)

	f, err := os.Open(srcFile)
	// Ensure the open file is always closed, ignoring any errors during the close
	defer func(f *os.File) {
		if f != nil {
			_ = f.Close()
		}
	}(f)

	if err != nil {
		return err
	}

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	h, err := zip.FileInfoHeader(stat)
	if err != nil {
		return err
	}
	h.Modified = ModifiedTimestamp
	h.Method = compression
	h.Name = srcFile
	h.Extra = extra

	// If dealing with a directory, we must ensure the h.Name field ends with a `/`
	if stat.IsDir() {
		if !strings.HasSuffix(h.Name, "/") {
			h.Name += "/"
		}
	}

	fw, err := zipWriter.CreateHeader(h)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		if _, err := io.Copy(fw, f); err != nil {
			return err
		}
	}
	// explicitly ignore any errors during the flush
	_ = zipWriter.Flush()

	return nil
}
