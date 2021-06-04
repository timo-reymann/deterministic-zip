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
)

const extension = ".zip"

func createFileName(input string) string {
	if strings.HasSuffix(input, extension) {
		return input
	}
	return input + extension
}

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
	if err != nil {
		return err
	}

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	header := zip.FileHeader{
		Name:   srcFile,
		Method: compression,
	}
	fw, err := zipWriter.CreateHeader(&header)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		if _, err := io.Copy(fw, f); err != nil {
			return err
		}
	}
	zipWriter.Flush()

	return nil
}
