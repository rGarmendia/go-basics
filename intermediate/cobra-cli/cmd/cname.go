package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cnameCmd)
}

var cnameCmd = &cobra.Command{
	Use:   "cname",
	Short: "Looks Up the CNAME for a particular Host",
	RunE: func(cmd *cobra.Command, args []string) error {
		// a simple lookup function
		cname, err := net.LookupCNAME(Host)
		if err != nil {
			return err
		}

		fmt.Println(cname)
		return nil
	},
}
