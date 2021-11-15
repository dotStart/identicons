package http

import (
	"fmt"
	"github.com/dotstart/identicons/library/identicons/icon"
	"net/http"
)

func NewInputHandlerFunc(gen icon.Generator) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		inputStr, ok := ctx.Value("input").(string)
		if !ok {
			http.Error(res, "no input given", http.StatusBadRequest)
			return
		}

		res.Header().Set("Content-Type", icon.ContentType)
		res.WriteHeader(http.StatusOK)
		gen.Write([]byte(inputStr), res)
	}
}

func NewInputRegistryHandlerFunc(registry *icon.Registry) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		genId, ok := ctx.Value("genId").(string)
		if !ok {
			http.Error(res, "no generator given", http.StatusBadRequest)
			return
		}

		inputStr, ok := ctx.Value("input").(string)
		if !ok {
			http.Error(res, "no input given", http.StatusBadRequest)
			return
		}

		gen := registry.Get(genId)
		if gen == nil {
			http.Error(res, fmt.Sprintf("no such generator: %s", genId), http.StatusNotFound)
			return
		}

		res.Header().Set("Content-Type", icon.ContentType)
		res.WriteHeader(http.StatusOK)
		gen.Write([]byte(inputStr), res)
	}
}
