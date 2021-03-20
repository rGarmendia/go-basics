package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	Host string

	rootCmd = &cobra.Command{
		Use:   "netctl",
		Short: "Website Lookup CLI",
		Long:  `Let's you query IPs, CNAMEs, MX records and Name Servers!`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVarP(&Host, "host", "", "", "Host to work")

}
