/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "xsops",
	Version: Version,
	Short:   "Use sops and json as a local secret store",
	Long: `xsops is a tool that allows you to use sops and json as a local secret store
It provides a simple interface to manage secrets securely and efficiently using sops. It 
defaults to using age for encryption.
	
For commands that use a URI, you can use the following formats:
- uri: sops:///path/to/secrets.json
- uri: file:///path/to/secrets.json
- path: ./xsops.secrets.json
- path: /path/to/secrets.json
- path: /home/user/.config/xsops/xsops.secrets.json
- path: .  (defaults to ./xsops.secrets.json in the current directory)
- special name: default (defaults to XDG_CONFIG_HOME/xsops/xsops.secrets.json)
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
