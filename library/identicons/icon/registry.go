package icon

// Registry encapsulates a set of known generators.
type Registry struct {
	ids        []string
	generators map[string]Generator
}

// EmptyRegistry creates a new registry devoid of any known generators.
func EmptyRegistry() *Registry {
	return &Registry{
		ids:        make([]string, 0),
		generators: make(map[string]Generator),
	}
}

// NewRegistry creates a registry with the given generator implementations.
func NewRegistry(generators ...Generator) *Registry {
	r := EmptyRegistry()
	for _, gen := range generators {
		r.Register(gen)
	}
	return r
}

// Register introduces a new generator to this registry.
//
// If a generator with the given identifier is already present within the registry, it will be
// replaced with the passed value.
func (r *Registry) Register(generator Generator) {
	_, exists := r.generators[generator.Id()]
	r.generators[generator.Id()] = generator

	if !exists {
		r.ids = append(r.ids, generator.Id())
	}
}

// Ids returns a list of known generator identifiers within this registry.
func (r *Registry) Ids() []string {
	return r.ids
}

// Get retrieves a given generator from this registry.
func (r *Registry) Get(id string) Generator {
	gen, ok := r.generators[id]
	if !ok {
		return nil
	}

	return gen
}
