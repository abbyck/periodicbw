package commands

import (
	"fmt"
	"os"

	"github.com/abbyck/periodicbw/runner"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "periodicbw",
	Short: "A small program to run speedtest CLI perioidically",
	Long:  `A small program to run SpeedTest CLI periodically and store the results onto a CSV file`,
	Run: func(cmd *cobra.Command, args []string) {
		runner.Runner()
	},
}

var Source string

func Execute() {
	// rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
