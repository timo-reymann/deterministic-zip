package zip

import (
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
		config cli.Configuration
		sha256 string
	}{
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/file.txt",
				},
			},
			sha256: "39366b547a236b1a82d08e6d3f1d351b0fae47263f2addaf579454c837f85bc5",
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder",
				},
			},
			sha256: "481cb3ec2e311b4112b90c37f1581bca8de65ebcf6987dcc610b098c209ecf51",
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
				},
			},
			sha256: "a13393b138cd0f6092b7efc69600f76cb0c38f4946291dd88bbf17e9b0edd562",
		},
		{
			config: cli.Configuration{
				SourceFiles: []string{
					"testdata/folder/file.txt",
					"testdata/file.txt",
				},
			},
			sha256: "2639b8988fff20364789d50ac13e7a656817ef65f86986d72c3909cf538dfbca",
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
			sha256: "18c98aed998d3693086ce0fc69f2e47731cac3307ef928172a7ad089d0743769",
		},
	}

	for _, tc := range testCases {
		rand.Shuffle(len(tc.config.SourceFiles), func(i, j int) {
			tc.config.SourceFiles[i], tc.config.SourceFiles[j] = tc.config.SourceFiles[j], tc.config.SourceFiles[i]
		})
		for i := 0; i < 20; i++ {
			// Create tempfile
			tc.config.ZipFile = createTmpFile() + extension

			_ = Create(&tc.config)

			sha256sum := checksum(tc.config.ZipFile)

			if tc.sha256 != sha256sum {
				t.Fatalf("Run #%d Expected checksum %s, but got %s, file: %s", i, tc.sha256, sha256sum, tc.config.ZipFile)
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
