package features

import "github.com/timo-reymann/deterministic-zip/pkg/cli"

type Feature interface {
	IsEnabled(config *cli.Configuration) bool
	Execute(config *cli.Configuration) error
	DebugName() string
}

var features = make([]Feature, 0)

func Features() *[]Feature {
	return &features
}

// register given feature, make sure to call in order they should be executed,
// it works like a pipeline!
func register(feature Feature) {
	features = append(features, feature)
}

func init() {
	register(Verbose{})
	register(Recursive{})
	register(Exclude{})
	register(NoDirectories{})
	register(Include{})
}
