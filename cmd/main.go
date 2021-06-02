package cmd

import (
	"fmt"
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"log"
	"os"
)

func Execute() {
	config := cli.NewConfiguration()
	err := config.Parse(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO Glue together files
	// TODO Zip together everything
	// TODO Add verbose output possibility

	fmt.Printf("%v", config)
}
