package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"math/rand"
	"os"
	"time"
)

type generateCommand struct {
	generator string
	random    bool
}

func GenerateCommand() subcommands.Command {
	return &generateCommand{}
}

func (*generateCommand) Name() string {
	return "generate"
}

func (*generateCommand) Synopsis() string {
	return "generates an identicon"
}

func (*generateCommand) Usage() string {
	generators := getGeneratorList()

	return `generate [options] <out-file> <input-string>

Generates an identicon and stores it in a given output file. For instance:

  $ identicons generate avatar.svg iliketrains

The content of the identicon is dictated by the given input string (in this case "iliketrains") and
will thus be the same for the same combination of generator and input string. Alternatively, a 
random icon may be generated by passing the -random option: 

  $ identicons generate avatar.svg -random

If desired, a different generator may be chosen:

  $ identicons generate -generator=classic avatar.svg iliketrains

The following generators are currently available:

` + generators + `

The following command line options are made available by this command:

`
}

func (cmd *generateCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.generator, "generator", defaultGenerator, "specifies the desired identicon generator")
	f.BoolVar(&cmd.random, "random", false, "generates a random identicon")
}

func (cmd *generateCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if !cmd.random && f.NArg() != 2 {
		_, _ = fmt.Fprintln(os.Stderr, "usage: identicons generator <out-file> <input-string>")
		return subcommands.ExitUsageError
	}
	if cmd.random && f.NArg() != 1 {
		_, _ = fmt.Fprintln(os.Stderr, "usage: identicons generator -random <out-file>")
		return subcommands.ExitUsageError
	}

	gen := generatorRegistry.Get(cmd.generator)
	if gen == nil {
		_, _ = fmt.Fprintf(os.Stderr, "no such generator: %s\n", cmd.generator)
		return subcommands.ExitUsageError
	}

	fileName := f.Arg(0)

	var input []byte
	if !cmd.random {
		input = []byte(f.Arg(1))
	} else {
		rand.Seed(time.Now().UnixNano())

		input = make([]byte, 32)
		_, _ = rand.Read(input)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE^os.O_TRUNC^os.O_WRONLY, 0644)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot open file \"%s\": %s\n", fileName, err)
		return subcommands.ExitFailure
	}
	defer func() { _ = file.Close() }()

	gen.Write(input, file)

	fmt.Println("done")
	return subcommands.ExitSuccess
}
