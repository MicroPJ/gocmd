package gocmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocmd",
	Short: "gocmd - Template GOLang project for CLI utilities",
	Long:  `Use this template to get up and running quickly`,
	PreRun: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd Run with args: %v\n", args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
