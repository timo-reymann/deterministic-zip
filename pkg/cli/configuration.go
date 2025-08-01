package cli

import (
	"errors"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// ModifiedTimestamp contains the default modification timestamp used for all files in the archive
var DefaultModifiedTimestamp = time.Date(2018, 11, 01, 0, 0, 0, 0, time.UTC)

// ErrMinimalParamsMissing states that the minimal arguments for the tool are not present, making it unprocessable
var ErrMinimalParamsMissing = errors.New("required arguments for target file and at least one source missing")

// Configuration represents the config for the cli and may be mutated by features
type Configuration struct {
	// ZipFile is the target zip file name
	ZipFile string

	// SourceFiles contains a flat list with paths that are either fully qualified or based on the pwd
	SourceFiles []string

	// Verbose outputs very detailed information
	Verbose bool

	// Recursive includes all child folders automatically
	Recursive bool

	// Directories includes directory entries in the zip file
	Directories bool

	// Exclude file patterns from the archive
	Exclude []string

	// Include patterns
	Include []string

	// CompressionMethod to use for the zip archive
	CompressionMethod string

	// Quiet eliminates all outputs
	Quiet bool

	// LogFilePath contains the path to the log
	LogFilePath string

	// LogFileAppend specifies if the file should be appended
	LogFileAppend bool

	// ModifiedDate represents the timestamp set for the modification time of all files in the zip file
	modifiedDate time.Time

	flagSet *flag.FlagSet
}

func (conf *Configuration) addBoolFlag(field *bool, long string, short string, val bool, usage string) {
	conf.flagSet.BoolVarP(field, long, short, val, usage)
}

func (conf *Configuration) addStringsFlag(field *[]string, long string, short string, val []string, usage string) {
	if short == "" {
		conf.flagSet.StringSliceVar(field, long, val, usage)
	} else {
		conf.flagSet.StringSliceVarP(field, long, short, val, usage)
	}
}

func (conf *Configuration) addStringFlag(field *string, long string, short string, val string, usage string) {
	conf.flagSet.StringVarP(field, long, short, val, usage)
}

func (conf *Configuration) defineFlags() {
	conf.flagSet = flag.NewFlagSet("deterministic-zip", flag.ErrorHandling(flag.ContinueOnError))
	conf.addBoolFlag(&conf.Verbose, "verbose", "v", false, "Verbose mode or print diagnostic version info.")
	conf.addBoolFlag(&conf.Directories, "directories", "D", false, "Include directories in the zip file.")
	conf.addBoolFlag(&conf.Recursive, "recurse-paths", "r", false, "Include all files verbose")
	conf.addBoolFlag(&conf.Quiet, "quiet", "q", false, "Quiet mode; eliminate informational messages")
	conf.addStringsFlag(&conf.Exclude, "exclude", "x", []string{}, "Exclude specific file patterns")
	conf.addStringsFlag(&conf.Include, "include", "i", []string{}, "Include only the specified file pattern")
	conf.addStringFlag(&conf.CompressionMethod, "compression-method", "Z", "deflate", "Set the default compression method. \nCurrently the main methods supported by zip are store and deflate. \nCompression method can be set to:\n\nstore       Setting the compression method to store forces to store entries with no compression. \n            This is generally faster than compressing entries, but results in no space savings.\n\ndeflate     This is the default method for zip. If zip determines that storing is better than deflation, the entry will be stored instead.\n")
	conf.addStringFlag(&conf.LogFilePath, "logfile-path", "", "", "Open a logfile at the given path.\nBy default any existing file at that location is overwritten, but the --log-append option will result in an existing file being opened and the new log information appended to any existing information.")
	conf.addBoolFlag(&conf.LogFileAppend, "log-append", "", false, "Append to existing logfile. Default is to overwrite.")
}

func (conf *Configuration) parseVarargs() error {
	remaining := conf.flagSet.Args()
	if len(remaining) < 2 {
		return ErrMinimalParamsMissing
	}

	conf.ZipFile = remaining[0]

	conf.SourceFiles = remaining[1:]

	return nil
}

func cleanPath(path string, isRetardedPlatform bool) string {
	if !isRetardedPlatform {
		return filepath.Clean(path)
	}

	windowsPath := filepath.Clean(path)
	properPath := strings.ReplaceAll(windowsPath, "\\", "/")
	return properPath
}

const isRetardedPlatform = runtime.GOOS == "windows"

// CleanPaths ensure that all directories and files in the file set have clean path names
func (conf *Configuration) CleanPaths() {
	cleaned := make([]string, 0, len(conf.SourceFiles))
	for _, f := range conf.SourceFiles {
		cleaned = append(cleaned, cleanPath(f, isRetardedPlatform))
	}
	conf.SourceFiles = cleaned
}

func (conf *Configuration) Help() {
	PrintCompactInfo()
	println("deterministic-zip [-options] [zipfile list]")
	conf.flagSet.PrintDefaults()
}

func (conf *Configuration) parseModifiedDate() (*time.Time, error) {
	sourceDateEpoch := os.Getenv("SOURCE_DATE_EPOCH")
	if sourceDateEpoch == "" {
		return &DefaultModifiedTimestamp, nil
	}

	sde, err := strconv.ParseInt(sourceDateEpoch, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid SOURCE_DATE_EPOCH: %s", err)
	}
	unixDate := time.Unix(sde, 0).UTC()
	return &unixDate, nil
}

// Parse the configuration from cli args
func (conf *Configuration) Parse() error {
	conf.defineFlags()

	isVersion := conf.flagSet.Bool("version", false, "Show version info")
	isHelp := conf.flagSet.BoolP("Help", "h", false, "Show available commands")
	err := conf.flagSet.Parse(os.Args[1:])

	if errors.Is(err, flag.ErrHelp) || (isHelp != nil && *isHelp) {
		conf.Help()
		return ErrAbort
	} else if *isVersion {
		PrintVersionInfo()
		return ErrAbort
	}

	modifiedDate, err := conf.parseModifiedDate()
	if err != nil {
		return err
	}
	conf.modifiedDate = *modifiedDate

	return conf.parseVarargs()
}

func (conf *Configuration) ModifiedDate() time.Time {
	if conf.modifiedDate.IsZero() {
		return DefaultModifiedTimestamp
	}
	return conf.modifiedDate
}

// NewConfiguration creates a new configuration
func NewConfiguration() *Configuration {
	return &Configuration{}
}
