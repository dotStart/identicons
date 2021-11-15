package http

import (
	"fmt"
	"github.com/dotstart/identicons/library/identicons/icon"
	"net"
	"net/http"
)

func NewRemoteAddressHandlerFunc(gen icon.Generator) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		addr, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			http.Error(res, "invalid remote address", http.StatusBadRequest)
			return
		}

		res.Header().Set("Content-Type", icon.ContentType)
		res.WriteHeader(http.StatusOK)
		gen.Write([]byte(addr), res)
	}
}

func NewRemoteAddressRegistryHandlerFunc(registry *icon.Registry) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		addr, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			http.Error(res, "invalid remote address", http.StatusBadRequest)
			return
		}

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
		gen.Write([]byte(addr), res)
	}
}
