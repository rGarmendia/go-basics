package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nsCmd)
}

var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Looks Up the NameServer for a particular Host",
	RunE: func(cmd *cobra.Command, args []string) error {
		// a simple lookup function
		ns, err := net.LookupNS(Host)
		if err != nil {
			return err
		}

		// we log the results to our console
		// using a trusty fmt.Println statement

		for i := 0; i < len(ns); i++ {
			fmt.Println(ns[i].Host)
		}
		return nil
	},
}
