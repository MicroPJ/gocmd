package gocmd

import (
	"fmt"

	"github.com/micropj/gocmd/pkg/gocmd"
	"github.com/spf13/cobra"
)

var oneCmd = &cobra.Command{
	Use:     "one",
	Aliases: []string{"o"},
	Short:   "Option one",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		res := gocmd.One(" ")
		fmt.Printf(res)
	},
}

func init() {
	rootCmd.AddCommand(oneCmd)
}
