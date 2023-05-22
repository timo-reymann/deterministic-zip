package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/fileset"
	"github.com/timo-reymann/deterministic-zip/pkg/features/filter"
	"github.com/timo-reymann/deterministic-zip/pkg/features/output"
)

// Feature is the high level abstraction of modular
type Feature interface {
	// IsEnabled checks if the feature should be executed
	IsEnabled(config *cli.Configuration) bool

	// Execute the feature
	Execute(config *cli.Configuration) error

	// DebugName prints the debuggable name
	DebugName() string
}

var features = make([]Feature, 0)

// Features returns all registered
func Features() *[]Feature {
	return &features
}

// register given feature, make sure to call in order they should be executed,
// it works like a pipeline!
func register(feature Feature) {
	features = append(features, feature)
}

func init() {
	// ATTENTION: Order is important -> its like a staging/pipeline system
	register(output.LogFile{})
	register(output.Verbose{})
	register(output.Quiet{})
	register(fileset.Recursive{})
	register(filter.Exclude{})
	register(filter.Include{})
	// Directories must always be processed after the Recursive, Exclude, Include modules
	register(fileset.Directories{})
	register(fileset.IgnoreTargetZip{})
}
