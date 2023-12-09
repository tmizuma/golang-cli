package htmlCmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type HtmlCmd struct {
	capitalize bool
}

func (*HtmlCmd) Name() string     { return "print" }
func (*HtmlCmd) Synopsis() string { return "Print args to stdout." }
func (*HtmlCmd) Usage() string {
	return `print [-capitalize] <some text>:
	Print args to stdout.
  `
}

func (h *HtmlCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&h.capitalize, "capitalize", false, "capitalize output")
}

func (h *HtmlCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println()
	return subcommands.ExitSuccess
}
