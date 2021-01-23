package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	configcmd "github.com/stnrd/doqu-cli/pkg/cmd/config"
	versioncmd "github.com/stnrd/doqu-cli/pkg/cmd/version"
	"github.com/stnrd/doqu-cli/pkg/cmdutil"
)

func NewCmdRoot(f *cmdutil.Factory, version, buildDate string) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "doqu <command> <subcommand> [flags]",
		Short: "DoQu CLI",
		Long:  `Work seamlessly with the Doqu from the command line.`,

		SilenceErrors: true,
		SilenceUsage:  true,
		Example: heredoc.Doc(`
			$ doqu doc create
			$ doqu link create
			$ doqu sync -url http://example.doqu.com
		`),
		Annotations: map[string]string{
			"help:environment": heredoc.Doc(`
				See 'doqu help environment' for the list of supported environment variables.
			`),
		},
	}

	cmd.SetOut(f.IOStreams.Out)
	cmd.SetErr(f.IOStreams.ErrOut)

	cs := f.IOStreams.ColorScheme()

	helpHelper := func(command *cobra.Command, args []string) {
		rootHelpFunc(cs, command, args)
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")
	cmd.SetHelpFunc(helpHelper)
	cmd.SetUsageFunc(rootUsageFunc)
	// cmd.SetFlagErrorFunc(rootFlagErrorFunc) -- Flag error is an error raised in flag processing

	formattedVersion := versioncmd.Format(version, buildDate)
	cmd.SetVersionTemplate(formattedVersion)
	cmd.Version = formattedVersion
	cmd.Flags().Bool("version", false, "Show doqu version")

	// TODO: add here the needed commands
	cmd.AddCommand(versioncmd.NewCmdVersion(f, version, buildDate))
	cmd.AddCommand(configcmd.NewCmdConfig(f))

	return cmd
}
