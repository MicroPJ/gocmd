package gocmd

import (
	"fmt"

	"github.com/micropj/gocmd/pkg/gocmd"
	"github.com/spf13/cobra"
)

var vsamCmd = &cobra.Command{
	Use:     "vsam",
	Aliases: []string{"v"},
	Short:   "Deploy BankDemo VSAM",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		res := gocmd.Vsam(" ")
		fmt.Printf(res)
	},
}

func init() {
	rootCmd.AddCommand(vsamCmd)
}
