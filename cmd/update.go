package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [args]",
	Short: "update the given odoo model",
	Long:  `This command update the needed api and type files for the given odoo model`,
	Example: `
	./godoo-cli update account.analytic.account
	./godoo-cli update account.analytic.account account.invoice
	./godoo-cli update all`,
	Run: upsert,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringP("endpoint", "e", "http://localhost:8069", "the odoo instance URI")
	updateCmd.PersistentFlags().StringP("database", "d", "database", "the name of the postgres database linked to odoo")
	updateCmd.PersistentFlags().StringP("username", "u", "admin", "the odoo instance administrator")
	updateCmd.PersistentFlags().StringP("password", "p", "password", "the odoo instance administrator password")
}
