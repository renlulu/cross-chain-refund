package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "cross-chain-refund",
	Short: "To run zilliqa refund server",
	Long:  `To run zilliqa refund server`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
