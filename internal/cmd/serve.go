package cmd

import (
	"context"
	"flag"
	"fmt"
	iconhttp "github.com/dotstart/identicons/library/identicons/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/subcommands"
	template "github.com/hashicorp/go-sockaddr/template"
	"net/http"
	"os"
	"time"
)

type serveCommand struct {
	bind      string
	context   string
	generator string
	timeout   time.Duration
}

func ServeCommand() subcommands.Command {
	return &serveCommand{}
}

func (*serveCommand) Name() string {
	return "serve"
}

func (*serveCommand) Synopsis() string {
	return "serves identicon requests via HTTP"
}

func (*serveCommand) Usage() string {
	generators := getGeneratorList()

	return `serve [options]

Launches a new HTTP server which generates identicons based on simple requests:

  $ identicons serve

If desired, a bind address may be specified to limit the server to a given network interface or
alter its bind port:

  $ identicons serve -bind=127.0.0.1:8080

Additionally, the server permits the central selection of a default generator via the generator
switch:

  $ identicons serve -generator=classic

Which will cause unqualified routes to use the classic generator. The following generators are 
available:

` + generators + `

This server implementation currently exposes the following routes:

  * /{input}         - icon for input, uses default generator
  * /random          - random icon, uses default generator
  * /ip              - icon for remote address, uses default generator
  * /{genId}/{input} - icon for input
  * /{genId}/random  - random icon
  * /{genId}/ip      - icon for remote address

The context path at which these endpoints are made available may be changed via the ctx flag:

  $ identicons serve -ctx=/foo

Which will relocate the endpoints to /foo (e.g. /random becomes /foo/random, /ip becomes /foo/ip, 
and so on).

The following command line options are provided by this command:

`
}

func (cmd *serveCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.bind, "bind", "0.0.0.0:8080", "selects the network address to bind to")
	f.StringVar(&cmd.context, "ctx", "", "selects the context path")
	f.StringVar(&cmd.generator, "generator", defaultGenerator, "selects the default generator")
	f.DurationVar(&cmd.timeout, "timeout", 15*time.Second, "selects the response timeout")
}

func (cmd *serveCommand) Execute(context.Context, *flag.FlagSet, ...interface{}) subcommands.ExitStatus {
	bind, err := template.Parse(cmd.bind)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "illegal bind address: %s\n", err)
		return subcommands.ExitUsageError
	}

	defaultGen := generatorRegistry.Get(cmd.generator)
	if defaultGen == nil {
		_, _ = fmt.Fprintf(os.Stderr, "no such generator: %s\n", cmd.generator)
		return subcommands.ExitUsageError
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(cmd.timeout))

	r.With(exposeParams).Get("/{genId}/{input}", iconhttp.NewInputRegistryHandlerFunc(generatorRegistry))
	r.With(exposeParams).Get("/{genId}/random", iconhttp.NewRandomRegistryHandlerFunc(generatorRegistry))
	r.With(exposeParams).Get("/{genId}/ip", iconhttp.NewRemoteAddressRegistryHandlerFunc(generatorRegistry))

	r.With(exposeParams).Get("/{input}", iconhttp.NewInputHandlerFunc(defaultGen))
	r.With(exposeParams).Get("/random", iconhttp.NewRandomHandlerFunc(defaultGen))
	r.With(exposeParams).Get("/ip", iconhttp.NewRemoteAddressHandlerFunc(defaultGen))

	if cmd.context != "" {
		wrapped := chi.NewRouter()
		wrapped.Mount(cmd.context, wrapped)
		r = wrapped
	}

	fmt.Printf("listening on %s\n\n", bind)
	err = http.ListenAndServe(bind, r)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to bind: %s\n", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func exposeParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		if r := chi.RouteContext(ctx); r != nil {
			for i, param := range r.URLParams.Keys {
				value := r.URLParams.Values[i]

				ctx = context.WithValue(ctx, param, value)
			}
		}

		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
