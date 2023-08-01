package cmd

import (
	"github.com/ives22/stool/apps/host"
	"github.com/spf13/cobra"
)

var userHome string

var KeyGenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "生成ssh密钥",
	Long:  "生成ssh密钥",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := host.NewSSHSetup()
		server.Home = userHome
		server.KeyGen()

		return nil
	},
}

func init() {
	KeyGenCmd.PersistentFlags().StringVarP(&userHome, "home", "H", "~", "用户家目录，echo $HOME")
	RootCmd.AddCommand(KeyGenCmd)
}