package delCmd

import (
	"context"
	"flag"
	"io/fs"
	"os"

	"github.com/google/subcommands"
)

type DelCmd struct {
	path string
}

func (*DelCmd) Name() string     { return "del" }
func (*DelCmd) Synopsis() string { return "Delete folders and files." }
func (*DelCmd) Usage() string {
	return `del [-path] <path>`
}

type OSFileOperator struct{}

func (o OSFileOperator) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (o OSFileOperator) Remove(path string) error {
	return os.Remove(path)
}

func (o OSFileOperator) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (o OSFileOperator) Mkdir(path string) error {
	return os.Mkdir(path, 0755)
}

func (d *DelCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&d.path, "path", "", "delete path")
}

func (d *DelCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	path := d.path
	fileOperator := OSFileOperator{}
	err := del(fileOperator, path)
	if err != nil {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
