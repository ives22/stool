package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "stool",
	Short: "stool shell工具",
	Long:  "stool shell工具",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no flags find")
	},
}