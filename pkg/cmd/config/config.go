package configcmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stnrd/doqu-cli/internal/config"
	"github.com/stnrd/doqu-cli/pkg/cmdutil"
)

func NewCmdConfig(f *cmdutil.Factory) *cobra.Command {
	longDoc := strings.Builder{}
	longDoc.WriteString("Display or change configuration settings for doqu.\n\n")
	longDoc.WriteString("Current respected settings:\n")
	for _, co := range config.ConfigOptions() {
		longDoc.WriteString(fmt.Sprintf("- %s: %s", co.Key, co.Description))
		if co.DefaultValue != "" {
			longDoc.WriteString(fmt.Sprintf(" (default: %q)", co.DefaultValue))
		}
		longDoc.WriteRune('\n')
	}

	cmd := &cobra.Command{
		Use:   "config <command>",
		Short: "Manage configuration for doqu",
		Long:  longDoc.String(),
	}

	cmd.AddCommand(newCmdConfigGet(f, nil))
	cmd.AddCommand(newCmdConfigSet(f, nil))

	return cmd
}
