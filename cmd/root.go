package cmd

import (
	// 3rd party
	"github.com/brutalgg/cli"
	"github.com/spf13/cobra"
	// internal
	"github.com/brutalgg/watchdog/pkg/banner"
	"github.com/brutalgg/watchdog/pkg/watchdog"
)

var rootCmd = &cobra.Command{
	Use:              "watchdog",
	Version:          "beta",
	PersistentPreRun: preChecks,
	Run:              run,
}

func init() {
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "Include verbose messages from program execution [error, warn, info, debug]")
	rootCmd.PersistentFlags().StringP("monitor", "m", "/Library/Group Containers/K36BKF7T3D.group.com.apple.configurator/Library/Caches/", "Location to monitor for cached IPA files")
	rootCmd.PersistentFlags().IntP("delay", "d", 2000, "Time in milliseconds for polling the filesystem")
	rootCmd.PersistentFlags().BoolP("banner", "q", false, "Disables banner")
	rootCmd.PersistentFlags().StringP("out", "o", "/Desktop/watchdog", "Location in the user's home directory for dumping files")
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
	m, _ := cmd.Flags().GetString("monitor")
	o, _ := cmd.Flags().GetString("out")
	d, _ := cmd.Flags().GetInt("delay")

	cli.Infoln("Initalizing watchdog instance")
	dog := watchdog.New(m, o, d)
	dog.Watch()
}
