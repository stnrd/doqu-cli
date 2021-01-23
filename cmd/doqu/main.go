package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"github.com/stnrd/doqu-cli/internal/build"
	"github.com/stnrd/doqu-cli/internal/update"
	"github.com/stnrd/doqu-cli/pkg/cmd/root"
	"github.com/stnrd/doqu-cli/pkg/cmdutil"
)

func main() {
	buildDate := build.Date
	buildVersion := build.Version

	updateMessageChan := make(chan *update.ReleaseInfo)
	go func() {
		rel, _ := checkForUpdate(buildVersion)
		updateMessageChan <- rel
	}()

	hasDebug := true
	// hasDebug := os.Getenv("DEBUG") != ""

	// Enable running gh from Windows File Explorer's address bar. Without this, the user is told to stop and run from a
	// terminal. With this, a user can clone a repo (or take other actions) directly from explorer.
	if len(os.Args) > 1 && os.Args[1] != "" {
		cobra.MousetrapHelpText = ""
	}

	cmdFactory := cmdutil.NewFactory(buildVersion)
	stderr := cmdFactory.IOStreams.ErrOut

	// TODO get something for the CmdFactory where we can access config / http client etc.
	rootCmd := root.NewCmdRoot(cmdFactory, buildVersion, buildDate)

	// Get the config
	cfg, err := cmdFactory.Config()
	if err != nil {
		fmt.Fprintf(stderr, "failed to read configuration:  %s\n", err)
		os.Exit(2)
	}

	if prompt, _ := cfg.Get("", "prompt"); prompt == "disabled" {
		cmdFactory.IOStreams.SetNeverPrompt(true)
	}

	if pager, _ := cfg.Get("", "pager"); pager != "" {
		cmdFactory.IOStreams.SetPager(pager)
	}

	expandedArgs := []string{}
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}

	cmd, _, err := rootCmd.Traverse(expandedArgs)
	if err != nil || cmd == rootCmd {
		// TODO: Figure out what's happening here!!
		// originalArgs := expandedArgs
		// isShell := false

		// expandedArgs, isShell, err = expand.ExpandAlias(cfg, os.Args, nil)
		// if err != nil {
		// 	fmt.Fprintf(stderr, "failed to process aliases:  %s\n", err)
		// 	os.Exit(2)
		// }

		// if hasDebug {
		// 	fmt.Fprintf(stderr, "%v -> %v\n", originalArgs, expandedArgs)
		// }

	}

	// cs := cmdFactory.IOStreams.ColorScheme()

	rootCmd.SetArgs(expandedArgs)

	if cmd, err := rootCmd.ExecuteC(); err != nil {
		printError(stderr, err, cmd, hasDebug)
		os.Exit(1)
	}
	if root.HasFailed() {
		os.Exit(1)
	}

	newRelease := <-updateMessageChan
	if newRelease != nil {
		msg := fmt.Sprintf("%s %s → %s\n%s",
			ansi.Color("A new release of doqu is available:", "yellow"),
			ansi.Color(buildVersion, "cyan"),
			ansi.Color(newRelease.Version, "cyan"),
			ansi.Color(newRelease.URL, "yellow"))

		fmt.Fprintf(stderr, "\n\n%s\n\n", msg)
	}

}

func printError(out io.Writer, err error, cmd *cobra.Command, debug bool) {
	if err == cmdutil.SilentError {
		return
	}

	fmt.Fprintln(out, err)

	var flagError *cmdutil.FlagError
	if errors.As(err, &flagError) || strings.HasPrefix(err.Error(), "unknown command ") {
		if !strings.HasSuffix(err.Error(), "\n") {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out, cmd.UsageString())
	}
}

// TODO: Implement later on an update check
func checkForUpdate(currentVersion string) (*update.ReleaseInfo, error) {
	// if !shouldCheckForUpdate() {
	// 	return nil, nil
	// }

	// client, err := basicClient(currentVersion)
	// if err != nil {
	// 	return nil, err
	// }

	// repo := updaterEnabled
	// stateFilePath := path.Join(config.ConfigDir(), "state.yml")
	// return update.CheckForUpdate(client, stateFilePath, repo, currentVersion)
	return &update.ReleaseInfo{
		Version: "v1.0.0",
		URL:     "https://www.google.nl",
	}, nil
}
