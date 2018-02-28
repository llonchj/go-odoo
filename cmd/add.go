package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [args]",
	Short: "add the given odoo model",
	Long:  `This command add the needed api and type files for the given odoo model`,
	Example: `
	./godoo-cli add account.analytic.account
	./godoo-cli add account.analytic.account account.invoice
	./godoo-cli add all`,
	Run: upsert,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringP("endpoint", "e", "http://localhost:8069", "the odoo instance URI")
	addCmd.PersistentFlags().StringP("database", "d", "database", "the name of the postgres database linked to odoo")
	addCmd.PersistentFlags().StringP("username", "u", "admin", "the odoo instance administrator")
	addCmd.PersistentFlags().StringP("password", "p", "password", "the odoo instance administrator password")
}
