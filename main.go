package main

import (
	"context"
	"flag"
	delCmd "golang-cli/delcmd"
	"golang-cli/formatcmd"
	"golang-cli/printcmd"
	"os"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&printcmd.PrintCmd{}, "")
	subcommands.Register(&formatcmd.FormatCmd{}, "")
	subcommands.Register(&delCmd.DelCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
