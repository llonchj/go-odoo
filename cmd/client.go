package cmd

import (
	"../api"
	"github.com/spf13/cobra"
)

func getClient(cmd *cobra.Command) (*api.Client, error) {
	c := &api.Client{}
	uri, err := cmd.PersistentFlags().GetString("endpoint")
	if err != nil {
		return nil, err
	}
	c, err = api.NewClient(uri, nil)
	if err != nil {
		return nil, err
	}
	db, err := cmd.PersistentFlags().GetString("database")
	if err != nil {
		return nil, err
	}
	admin, err := cmd.PersistentFlags().GetString("username")
	if err != nil {
		return nil, err
	}
	password, err := cmd.PersistentFlags().GetString("password")
	if err != nil {
		return nil, err
	}
	return c, c.Login(db, admin, password)
}
