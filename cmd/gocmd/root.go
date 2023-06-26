package gocmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocmd",
	Short: "gocmd - Template GOLang project for CLI utilities",
	Long:  `Run "gocmd -h" for help on any command`,
	PreRun: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Run 'gocmd -h' for help on any command\n")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
