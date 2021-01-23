package versioncmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stnrd/doqu-cli/pkg/cmdutil"
)

func NewCmdVersion(f *cmdutil.Factory, version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(f.IOStreams.Out, Format(version, buildDate))
		},
	}

	return cmd
}

func Format(version, buildDate string) string {
	version = strings.TrimPrefix(version, "v")

	var dateStr string
	if buildDate != "" {
		dateStr = fmt.Sprintf(" (%s)", buildDate)
	}

	return fmt.Sprintf("doqu cli version %s%s\n", version, dateStr)
}
