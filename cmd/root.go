package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godoo-cli COMMAND [ARGS]",
	Short: "Prepare your custom odoo api environment",
	Long:  "Tool to prepare environment for odoo golang api wrapper",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
