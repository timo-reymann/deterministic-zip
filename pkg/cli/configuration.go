package cli

import (
	"errors"
	"flag"
	"os"
)

// Configuration represents the config for the cli
type Configuration struct {
	Verbose     bool
	Recursive   bool
	ZipFile     string
	SourceFiles []string
}

func (conf *Configuration) addBoolFlag(field *bool, long string, short string, val bool, usage string) {
	flag.BoolVar(field, long, val, usage)
	flag.BoolVar(field, short, val, usage+" (short)")
}

// Parse the configuration from cli args
func (conf *Configuration) Parse(args []string) error {
	conf.addBoolFlag(&conf.Verbose, "verbose", "v", false, "verbose mode for debug outputs and trouble shooting")
	conf.addBoolFlag(&conf.Recursive, "recursive", "R", false, "include all files verbose")

	cli := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if err := cli.Parse(args); err != nil {
		return err
	}

	remaining := flag.Args()
	if len(remaining) < 2 {
		return errors.New("specify at least the destination package and source files")
	}

	conf.ZipFile = remaining[0]
	conf.SourceFiles = remaining[1:]

	return nil
}

// NewConfiguration creates a new configuration
func NewConfiguration() *Configuration {
	return &Configuration{}
}
