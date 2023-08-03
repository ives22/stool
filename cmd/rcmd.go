package cmd

import (
	"fmt"
	"github.com/ives22/stool/apps/host"
	"github.com/spf13/cobra"
)

var (
	rUser   string
	rPort   int64
	rIPFile string
	rCmd    string
)

var RunShellCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Executing shell commands on remote servers",
	Long:  "Executing shell commands on remote servers",
	RunE: func(cmd *cobra.Command, args []string) error {

		// 检查 pwd 参数是否为空
		if rCmd == "" {
			return fmt.Errorf("command is required")
		}

		// NewPush
		cli := host.NewClientsConf()

		if rUser == "" {
			rUser = cli.GetUser()
		}

		if rIPFile == "" {
			ipFile = "ip.txt"
		}

		// 初始化
		cli.InitClientForKey(rIPFile, rUser, rPort)

		// 执行命令
		cli.Command(rCmd)

		return nil
	},
}

func init() {
	RunShellCmd.PersistentFlags().StringVarP(&rUser, "username", "u", "", "user name, default current user")
	RunShellCmd.PersistentFlags().Int64VarP(&rPort, "port", "P", 22, "ssh port")
	RunShellCmd.PersistentFlags().StringVarP(&rIPFile, "file", "f", "ip.txt", "ip list file")
	RunShellCmd.PersistentFlags().StringVarP(&rCmd, "command", "c", "", "Command to execute (required)")

	RootCmd.AddCommand(RunShellCmd)
}