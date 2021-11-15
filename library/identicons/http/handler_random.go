package http

import (
	"fmt"
	"github.com/dotstart/identicons/library/identicons/icon"
	"math/rand"
	"net/http"
)

func NewRandomHandlerFunc(gen icon.Generator) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		input := make([]byte, 32)
		_, _ = rand.Read(input)

		res.Header().Set("Content-Type", icon.ContentType)
		res.WriteHeader(http.StatusOK)
		gen.Write(input, res)
	}
}

func NewRandomRegistryHandlerFunc(registry *icon.Registry) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		input := make([]byte, 32)
		_, _ = rand.Read(input)

		genId, ok := ctx.Value("genId").(string)
		if !ok {
			http.Error(res, "no generator given", http.StatusBadRequest)
			return
		}

		gen := registry.Get(genId)
		if gen == nil {
			http.Error(res, fmt.Sprintf("no such generator: %s", genId), http.StatusNotFound)
			return
		}

		res.Header().Set("Content-Type", icon.ContentType)
		res.WriteHeader(http.StatusOK)
		gen.Write(input, res)
	}
}
