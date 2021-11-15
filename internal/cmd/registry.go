package cmd

import "github.com/dotstart/identicons/library/identicons"

const defaultGenerator = "modern"

var generatorRegistry = identicons.DefaultRegistry()

func getGeneratorList() string {
	generators := ""
	for _, gen := range generatorRegistry.Ids() {
		if generators != "" {
			generators += "\n"
		}

		generators += " * " + gen

		if gen == defaultGenerator {
			generators += " (default)"
		}
	}
	return generators
}
