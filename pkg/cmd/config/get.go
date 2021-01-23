package configcmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/stnrd/doqu-cli/internal/config"
	"github.com/stnrd/doqu-cli/pkg/cmdutil"
	"github.com/stnrd/doqu-cli/pkg/iostreams"
)

type GetOptions struct {
	IO     *iostreams.IOStreams
	Config config.Config

	Hostname string
	Key      string
}

func newCmdConfigGet(f *cmdutil.Factory, runF func(*GetOptions) error) *cobra.Command {
	opts := &GetOptions{
		IO: f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "get <key>",
		Short: "Print the value of a given configuration key",
		Example: heredoc.Doc(`
			$ doqu config get editor
			vim
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := f.Config()
			if err != nil {
				return err
			}
			opts.Config = config
			opts.Key = args[0]

			if runF != nil {
				return runF(opts)
			}

			return getRun(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Hostname, "host", "h", "", "Get per-host setting")

	return cmd
}

func getRun(opts *GetOptions) error {
	val, err := opts.Config.Get("", opts.Key)
	if err != nil {
		return err
	}

	if val != "" {
		fmt.Fprintf(opts.IO.Out, "%s\n", val)
	}
	return nil
}
