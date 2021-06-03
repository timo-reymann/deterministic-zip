package features

import "github.com/timo-reymann/deterministic-zip/pkg/cli"

type Feature interface {
	IsEnabled(config *cli.Configuration) bool
	Execute(config *cli.Configuration)
}

var features = make([]Feature, 0)

func Features() *[]Feature {
	return &features
}

func register(feature Feature) {
	features = append(features, feature)
}

func init() {
	register(Verbose{})
}
