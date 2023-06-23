package ppp

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ppp",
	Short: "ppp - Poor Persons Pipeline, a simple CLI to deploy BankDemo",
	Long:  `One can use ppp to deploy BankDemo from the GitHub Public Repository`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
