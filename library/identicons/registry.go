package identicons

import (
	"github.com/dotstart/identicons/library/identicons/icon"
	"github.com/dotstart/identicons/library/identicons/icon/block"
	"github.com/dotstart/identicons/library/identicons/icon/circle"
	"github.com/dotstart/identicons/library/identicons/icon/circlematrix"
	"github.com/dotstart/identicons/library/identicons/icon/classic"
	"github.com/dotstart/identicons/library/identicons/icon/modern"
)

// DefaultRegistry creates a registry with all library-provided generators registered.
func DefaultRegistry() *icon.Registry {
	return icon.NewRegistry(
		block.New(),
		circle.New(),
		circlematrix.New(),
		classic.New(),
		modern.New(),
	)
}
