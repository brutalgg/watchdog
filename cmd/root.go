package cmd

import (
	// 3rd party
	"github.com/brutalgg/cli"
	"github.com/spf13/cobra"
	// internal
	"github.com/brutalgg/watchdog/pkg/banner"
)

var rootCmd = &cobra.Command{
	Use:              "watchdog",
	Version:          "beta",
	PersistentPreRun: preChecks,
	Run:              run,
}

func init() {
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "Include verbose messages from program execution [error, warn, info, debug]")
	rootCmd.PersistentFlags().StringP("monitor", "m", "/Library/Group Containers/K36BKF7T3D.group.com.apple.configurator/Library/Caches/", "Folder to monitor for cached IPA files")
	rootCmd.PersistentFlags().BoolP("banner", "q", false, "Disables banner")
	rootCmd.PersistentFlags().StringP("out", "o", "/tmp/ipawatchdog", "Location to dump detected IPA files")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cli.Fatalln(err)
	}
}

func preChecks(cmd *cobra.Command, args []string) {
	l, _ := cmd.Flags().GetString("loglevel")
	switch l {
	case "error":
		cli.SetPrintLevel(3)
	case "warn":
		cli.SetPrintLevel(2)
	case "debug":
		cli.SetPrintLevel(0)
	default:
		cli.SetPrintLevel(1)
	}

	b, _ := cmd.Flags().GetBool("banner")
	if !b {
		banner.Print()
	}

}

func run(cmd *cobra.Command, args []string) {
}
