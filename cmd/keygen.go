package cmd

import (
	"github.com/ives22/stool/apps/host"
	"github.com/spf13/cobra"
)

var userHome string

var KeyGenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Generating public/private rsa key pair for user",
	Long:  "Generating public/private rsa key pair for user",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := host.NewSSHSetup()
		server.Home = userHome
		// 创建ssh密钥
		server.KeyGen()

		return nil
	},
}

func init() {
	KeyGenCmd.PersistentFlags().StringVarP(&userHome, "home", "H", "~", "user home，echo $HOME")
	RootCmd.AddCommand(KeyGenCmd)
}