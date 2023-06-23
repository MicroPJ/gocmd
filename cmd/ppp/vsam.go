package ppp

import (
	"fmt"

	"github.com/micropj/ppp/pkg/ppp"
	"github.com/spf13/cobra"
)

var vsamCmd = &cobra.Command{
	Use:     "vsam",
	Aliases: []string{"v"},
	Short:   "Deploy BankDemo VSAM",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		res := ppp.Vsam(" ")
		fmt.Printf(res)
	},
}

func init() {
	rootCmd.AddCommand(vsamCmd)
}
