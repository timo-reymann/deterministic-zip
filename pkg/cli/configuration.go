package cli

import (
	"errors"
	flag "github.com/spf13/pflag"
)

// Configuration represents the config for the cli and may be mutated by features
type Configuration struct {
	// ZipFile contains the target zip file name
	ZipFile string
	// SourceFiles contains a flat list with paths that are either fully qualified or based on the pwd
	SourceFiles []string
	// Verbose states if the output should contain very detailed information
	Verbose bool
	// Recursive states if folders should be included recursively
	Recursive bool
	// Exclude contains file patterns to exclude from the archive
	Exclude []string
}

func (conf *Configuration) addBoolFlag(field *bool, long string, short string, val bool, usage string) {
	flag.BoolVarP(field, long, short, val, usage)
}

func (conf *Configuration) addStringsFlag(field *[]string, long string, short string, val []string, usage string) {
	if short == "" {
		flag.StringSliceVar(field, long, val, usage)
	} else {
		flag.StringSliceVarP(field, long, short, val, usage)
	}
}

// Parse the configuration from cli args
func (conf *Configuration) Parse() error {
	conf.addBoolFlag(&conf.Verbose, "verbose", "v", false, "verbose mode for debug outputs and trouble shooting")
	conf.addBoolFlag(&conf.Recursive, "recursive", "R", false, "include all files verbose")
	conf.addStringsFlag(&conf.Exclude, "exclude", "", []string{}, "exclude specific file pattern")

	flag.Parse()

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