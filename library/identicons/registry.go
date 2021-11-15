package identicons

import (
	"github.com/dotstart/identicons/library/identicons/icon"
	"github.com/dotstart/identicons/library/identicons/icon/circlematrix"
	"github.com/dotstart/identicons/library/identicons/icon/classic"
	"github.com/dotstart/identicons/library/identicons/icon/modern"
)

// DefaultRegistry creates a registry with all library-provided generators registered.
func DefaultRegistry() *icon.Registry {
	return icon.NewRegistry(
		circlematrix.New(),
		classic.New(),
		modern.New(),
	)
}
