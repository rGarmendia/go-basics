package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of netctl",
	Long:  `All software has versions. This is netctl's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("netctl v0.1")
	},
}
