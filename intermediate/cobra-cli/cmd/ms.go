package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mxCmd)
}

var mxCmd = &cobra.Command{
	Use:   "mx",
	Short: "Looks Up the MS for a particular Host",
	RunE: func(cmd *cobra.Command, args []string) error {
		// a simple lookup function
		mx, err := net.LookupMX(Host)
		if err != nil {
			return err
		}

		for i := 0; i < len(mx); i++ {
			fmt.Println(mx[i])
		}
		return nil
	},
}
