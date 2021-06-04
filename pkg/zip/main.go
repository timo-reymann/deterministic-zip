package zip

import (
	"archive/zip"
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
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

func Create(c *cli.Configuration) error {
	finalName := createFileName(c.ZipFile)
	newZipFile, err := os.OpenFile(finalName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	c.ZipFile = finalName

	zipWriter := zip.NewWriter(newZipFile)
	sort.Strings(c.SourceFiles)
	for _, srcFile := range c.SourceFiles {
		if err := appendFile(srcFile, zipWriter); err != nil {
			return err
		}
	}

	return zipWriter.Close()
}

func appendFile(srcFile string, zipWriter *zip.Writer) error {
	f, err := os.Open(srcFile)
	if err != nil {
		return err
	}

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	fw, err := zipWriter.CreateHeader(&zip.FileHeader{
		Name: srcFile,
	})
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		if _, err := io.Copy(fw, f); err != nil {
			return err
		}
	}
	return nil
}
