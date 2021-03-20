package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ipCmd)
}

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Looks Up the IP for a particular Host",
	RunE: func(cmd *cobra.Command, args []string) error {
		// a simple lookup function
		ip, err := net.LookupIP(Host)
		if err != nil {
			return err
		}

		for i := 0; i < len(ip); i++ {
			fmt.Println(ip[i])
		}
		return nil
	},
}
