package gocmd

import (
	"fmt"

	"github.com/micropj/gocmd/pkg/gocmd"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(oneCmd)
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GOCMD",
	Long:  `All software has versions. This is GOCMD's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GOCMD Template v0.1\n")
	},
}

var oneCmd = &cobra.Command{
	Use:     "one",
	Aliases: []string{"o"},
	Short:   "Option one",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		res := gocmd.One(args)
		fmt.Printf(res)
	},
}
