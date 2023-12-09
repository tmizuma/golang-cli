package formatcmd

import (
	"context"
	"flag"
	"log"

	"github.com/google/subcommands"
)

type FormatCmd struct {
	inputFolder  string
	outputFolder string
	chunkSize    int
}

func (*FormatCmd) Name() string     { return "format" }
func (*FormatCmd) Synopsis() string { return "Rename files and transfer to deploy folder." }
func (*FormatCmd) Usage() string {
	return `
	format [-inputFolder] <input folder> [-outputFolder] <output folder> [-chunkSize] <chunk size>
  `
}

func (p *FormatCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.inputFolder, "inputFolder", "", "input folder")
	f.StringVar(&p.outputFolder, "outputFolder", "", "output folder")
	f.IntVar(&p.chunkSize, "chunkSize", 0, "chunk size")
}

func (p *FormatCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fileNames, err := readFileNamesByPath(p.inputFolder)
	if err != nil {
		log.Fatal(err)
		return subcommands.ExitFailure
	}

	index, err := getMaxFolderIndex(p.outputFolder)
	if err != nil {
		log.Fatal(err)
		return subcommands.ExitFailure
	}

	fileTransferInfo, err := createNewPathInfo(fileNames, p.inputFolder, p.chunkSize, index+1)
	if err != nil {
		log.Fatal(err)
		return subcommands.ExitFailure
	}

	err = mvPath(fileTransferInfo, p.outputFolder)
	if err != nil {
		log.Fatal(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
