package zip

import (
	"archive/zip"
	"crypto/sha256"
	"fmt"
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"io"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func checksum(file string) string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

type expectedFile struct {
	name  string
	isDir bool
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		config          cli.Configuration
		sha256          string
		compression     uint16
		customExtension string
		zipFiles        []expectedFile
	}{
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/file.txt",
				},
			},
			sha256:      "32198c12721f7bc3b0fffad9df16c3e9fa56c4b698d5390f74dd1e7e74fbb915",
			compression: zip.Store,
			zipFiles: []expectedFile{
				{
					name: "testdata/file.txt",
				},
			},
		},
		{
			config: cli.Configuration{
				Directories: true,
				SourceFiles: []string{
					"testdata/folder",
				},
			},
			sha256:      "7839c2a3939b278e9b24e02621e6bdb07f4c32f79e111661ee9948f7516009c3",
			compression: zip.Store,
			zipFiles: []expectedFile{
				{
					name: "testdata/folder/",
				},
			},
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
				},
			},
			sha256:      "dd97707d68eda2563e0686e29934e4a7cd0437e761e9d02fdc6456cb3fd91eb7",
			compression: zip.Store,
			zipFiles: []expectedFile{
				{
					name: "testdata/folder/file.txt",
				},
			},
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
					"testdata/file.txt",
				},
			},
			sha256:      "b18ca34af3f15c04ec624e286412f44b2ed5c83e83d93b4b5b148aa03477ee9f",
			compression: zip.Store,
			zipFiles: []expectedFile{
				{
					name: "testdata/folder/file.txt",
				},
				{
					name: "testdata/file.txt",
				},
			},
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata",
					"testdata/file.txt",
					"testdata/folder",
					"testdata/folder/file.txt",
				},
			},
			sha256:      "b18ca34af3f15c04ec624e286412f44b2ed5c83e83d93b4b5b148aa03477ee9f",
			compression: zip.Store,
			zipFiles: []expectedFile{
				{
					name: "testdata/file.txt",
				},
				{
					name: "testdata/folder/file.txt",
				},
			},
		},
		{
			config: cli.Configuration{
				Directories: true,
				SourceFiles: []string{
					"testdata",
					"testdata/file.txt",
					"testdata/folder",
					"testdata/folder/file.txt",
				},
			},
			sha256:      "e2431c807ee3f202e84f66ed3756ae736eb890916cf1737420708bed2181c5e0",
			compression: zip.Store,
			zipFiles: []expectedFile{
				{
					name: "testdata/",
				},
				{
					name: "testdata/file.txt",
				},
				{
					name: "testdata/folder/",
				},
				{
					name: "testdata/folder/file.txt",
				},
			},
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata",
					"testdata/file.txt",
					"testdata/folder",
					"testdata/folder/file.txt",
				},
			},
			sha256:      "8b3eeacdd0c5c265a67bf465d9fc7d3ed0c041fc27534fb3f14b34d5a2b0b518",
			compression: zip.Deflate,
			zipFiles: []expectedFile{
				{
					name: "testdata/file.txt",
				},
				{
					name: "testdata/folder/file.txt",
				},
			},
		},
		{
			config: cli.Configuration{
				Directories: true,
				SourceFiles: []string{
					"testdata",
					"testdata/file.txt",
					"testdata/folder",
					"testdata/folder/file.txt",
				},
			},
			sha256:      "9501f16697415f9de62aab8e28925111abfac435842b455ea1bd36852a5b6adc",
			compression: zip.Deflate,
			zipFiles: []expectedFile{
				{
					name: "testdata/",
				},
				{
					name: "testdata/file.txt",
				},
				{
					name: "testdata/folder/",
				},
				{
					name: "testdata/folder/file.txt",
				},
			},
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
					"testdata/file.txt",
				},
			},
			sha256:      "8b3eeacdd0c5c265a67bf465d9fc7d3ed0c041fc27534fb3f14b34d5a2b0b518",
			compression: zip.Deflate,
			zipFiles: []expectedFile{
				{
					name: "testdata/file.txt",
				},
				{
					name: "testdata/folder/file.txt",
				},
			},
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
					"testdata/file.txt",
				},
			},
			sha256:      "8b3eeacdd0c5c265a67bf465d9fc7d3ed0c041fc27534fb3f14b34d5a2b0b518",
			compression: zip.Deflate,
			zipFiles: []expectedFile{
				{
					name: "testdata/file.txt",
				},
				{
					name: "testdata/folder/file.txt",
				},
			},
			customExtension: "rock",
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
					"testdata/file.txt",
				},
			},
			sha256:      "8b3eeacdd0c5c265a67bf465d9fc7d3ed0c041fc27534fb3f14b34d5a2b0b518",
			compression: zip.Deflate,
			zipFiles: []expectedFile{
				{
					name: "testdata/file.txt",
				},
				{
					name: "testdata/folder/file.txt",
				},
			},
			customExtension: "rock.zip",
		},
	}

	for idx, tc := range testCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			rand.Shuffle(len(tc.config.SourceFiles), func(i, j int) {
				tc.config.SourceFiles[i], tc.config.SourceFiles[j] = tc.config.SourceFiles[j], tc.config.SourceFiles[i]
			})
			for i := 0; i < 20; i++ {
				tempFile := createTmpFile()
				tempFileZip := tempFile + extension
				if tc.customExtension != "" {
					tempFileZip = tempFile + "." + tc.customExtension
				}
				// Create tempfile
				tc.config.ZipFile = tempFileZip

				_ = Create(&tc.config, tc.compression)

				sha256sum := checksum(tc.config.ZipFile)

				if tc.sha256 != sha256sum {
					t.Fatalf("Run #%d Expected checksum %s, but got %s, file: %s", i, tc.sha256, sha256sum, tc.config.ZipFile)
				}

				if tc.config.ZipFile != tempFileZip && tc.customExtension == "" {
					t.Fatalf("Expected final zip name to be either ending with .zip or %s, but got %s", extension, tc.config.ZipFile)
				}

				r, err := zip.OpenReader(tempFileZip)
				if err != nil {
					t.Fatal(err)
				}

				var foundFile *zip.File = nil

				if len(tc.zipFiles) == 0 && len(tc.zipFiles) != 0 {
					t.Fatalf("Expected no files in zip, but got %v", tc.zipFiles)
				}

				for _, expectedFile := range tc.zipFiles {
					foundFile = nil
					for _, file := range r.File {
						if expectedFile.name == file.Name {
							foundFile = file
							break
						}
					}

					if foundFile == nil {
						t.Fatalf("Expected file %s to be in archive", expectedFile.name)
					}

					if foundFile.Modified.Sub(ModifiedTimestamp) != 0 {
						t.Fatalf("Modified timestamp not reset for file %s", expectedFile.name)
					}

					if (foundFile.FileHeader.Mode() == os.ModeDir) != expectedFile.isDir {
						t.Fatalf("Expected file %s to be directory == %v", expectedFile.name, expectedFile.isDir)
					}
				}

				_ = r.Close()
				_ = os.Remove(tempFileZip)
			}
		})

	}
}

func createTmpFile() string {
	f, _ := os.CreateTemp(os.TempDir(), "go_test_")
	_ = f.Close()
	_ = os.Remove(f.Name())
	return f.Name()
}
