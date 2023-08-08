package cmd

import (
	"fmt"
	"github.com/ives22/stool/apps/host"
	"github.com/spf13/cobra"
)

var (
	user   string
	pwd    string
	port   int64
	ipFile string
)

var PushKeyCmd = &cobra.Command{
	Use:   "pushkey",
	Short: "SSH password free configuration, with the ip.txt file as the target host",
	Long:  "SSH password free configuration, with the ip.txt file as the target host",
	RunE: func(cmd *cobra.Command, args []string) error {

		// 检查 pwd 参数是否为空
		if pwd == "" {
			return fmt.Errorf("password is required")
		}

		// NewPush
		push := host.NewClientsConf()

		if user == "" {
			user = push.GetUser()
		}

		// 初始化
		push.Init(ipFile, user, pwd, port)

		// 密钥推送
		push.PushKey()

		return nil
	},
}

func init() {
	PushKeyCmd.PersistentFlags().StringVarP(&user, "username", "u", "", "user name default current user")
	PushKeyCmd.PersistentFlags().StringVarP(&pwd, "password", "p", "", "user password (required)")
	PushKeyCmd.PersistentFlags().Int64VarP(&port, "port", "P", 22, "ssh port")
	PushKeyCmd.PersistentFlags().StringVarP(&ipFile, "file", "f", "ip.txt", "ip list file")

	RootCmd.AddCommand(PushKeyCmd)
}