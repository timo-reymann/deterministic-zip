package cmd

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
	"github.com/timo-reymann/deterministic-zip/pkg/zip"
	"log"
	"os"
)

func errCheck(err error) {
	if err == cli.ErrAbort {
		os.Exit(0)
	}

	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
}

// Execute the application
func Execute() {
	config := cli.NewConfiguration()
	errCheck(config.Parse())

	compressionSpec, err := zip.GetCompressionMethod(config.CompressionMethod)
	errCheck(err)

	for _, f := range *features.Features() {
		if f.IsEnabled(config) {
			output.Debugf("Executing feature %s ...", f.DebugName())
			errCheck(f.Execute(config))
		}
	}

	output.Debugf("Using go zip compression method %d", compressionSpec)

	if err = zip.Create(config, compressionSpec); err != nil {
		log.Fatalln(err)
	}
}
