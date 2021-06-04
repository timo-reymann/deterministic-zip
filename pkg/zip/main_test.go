package zip

import (
	"archive/zip"
	"crypto/sha256"
	"fmt"
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"io"
	"math/rand"
	"os"
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

func TestCreate(t *testing.T) {
	testCases := []struct {
		config      cli.Configuration
		sha256      string
		compression uint16
	}{
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/file.txt",
				},
			},
			sha256:      "a899daf39908313e153c1ddc5676e22ce8fc84672e91f38d15113b09ca64f3c4",
			compression: zip.Store,
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder",
				},
			},
			sha256:      "7843fd78addc5f54432b1390978392dfddd5eb29e8aaf9151b5b3fdf87e2780b",
			compression: zip.Store,
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
				},
			},
			sha256:      "a5064a48c3ab57a6ac746cd09ce712e91b94d31aada5995a3d5cd02baf47e9d8",
			compression: zip.Store,
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
					"testdata/file.txt",
				},
			},
			sha256:      "d83d02e2820f8efffafcb4380dca07ee3b242ea2fdf663782c4519200ac2541c",
			compression: zip.Store,
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
			sha256:      "7ace043a37d5e62273c556f7c403f5d45a34400fb69a5776e3be1e153728e022",
			compression: zip.Store,
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
			sha256:      "c8e402dd0f5cad2fb568818d3b12e3745d785e9e8ca079737113a768d38c3b92",
			compression: zip.Deflate,
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
					"testdata/file.txt",
				},
			},
			sha256:      "dde82b1ef04349710e40bca6fd5f842fd11cf7b529d2d95ec6d6a87ebb0e9882",
			compression: zip.Deflate,
		},
	}

	for _, tc := range testCases {
		rand.Shuffle(len(tc.config.SourceFiles), func(i, j int) {
			tc.config.SourceFiles[i], tc.config.SourceFiles[j] = tc.config.SourceFiles[j], tc.config.SourceFiles[i]
		})
		for i := 0; i < 20; i++ {
			tempFile := createTmpFile()
			// Create tempfile
			tc.config.ZipFile = tempFile

			_ = Create(&tc.config, tc.compression)

			sha256sum := checksum(tc.config.ZipFile)

			if tc.sha256 != sha256sum {
				t.Fatalf("Run #%d Expected checksum %s, but got %s, file: %s", i, tc.sha256, sha256sum, tc.config.ZipFile)
			}

			if tc.config.ZipFile != tempFile+extension {
				t.Fatalf("Expected final zip name to be overriden")
			}
		}
	}
}

func createTmpFile() string {
	f, _ := os.CreateTemp(os.TempDir(), "go_test_")
	_ = f.Close()
	_ = os.Remove(f.Name())
	return f.Name()
}
