package cmd

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
	"github.com/timo-reymann/deterministic-zip/pkg/zip"
	"log"
	"os"
)

// Variable to allow mocking os.Exit in tests
var osExit = os.Exit

func errCheck(err error, c *cli.Configuration) {
	if err == cli.ErrAbort {
		osExit(0)
	}

	if err == cli.ErrMinimalParamsMissing {
		c.Help()
		osExit(2)
	}

	if err != nil {
		log.Println(err)
		osExit(2)
	}
}

// Execute the application
func Execute() {
	c := cli.NewConfiguration()
	errCheck(c.Parse(), c)

	compressionSpec, err := zip.GetCompressionMethod(c.CompressionMethod)
	errCheck(err, c)

	for _, f := range *features.Features() {
		// Clean paths before each feature to take that logic off the features
		c.CleanPaths()

		if f.IsEnabled(c) {
			output.Debugf("Executing feature %s ...", f.DebugName())
			errCheck(f.Execute(c), c)
		}
	}

	output.Debugf("Using go zip compression method %d", compressionSpec)

	if err = zip.Create(c, compressionSpec); err != nil {
		log.Fatalln(err)
	}
}
