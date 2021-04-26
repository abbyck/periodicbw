package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of periodicbw",
	Long:  `All software has versions. This is periodicBW's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PeriodicBW - Run Periodic Internet Bandwidth tests using speedtest CLI v0.0.1")
	},
}
