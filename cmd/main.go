package cmd

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features"
	"github.com/timo-reymann/deterministic-zip/pkg/zip"
	"log"
)

func Execute() {
	config := cli.NewConfiguration()
	err := config.Parse()
	if err == cli.ErrAbort {
		return
	} else if err != nil {
		log.Fatalln(err)
	}

	for _, f := range *features.Features() {
		if f.IsEnabled(config) {
			err := f.Execute(config)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	if err = zip.Create(config); err != nil {
		log.Fatalln(err)
	}
}
