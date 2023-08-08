package cmd

import (
	"fmt"
	"github.com/ives22/stool/apps/host"
	"github.com/spf13/cobra"
)

var (
	srcFile string
	dstFile string
	pIPFile string
)

var PushFile = &cobra.Command{
	Use:   "push",
	Short: "SSH password free configuration, with the ip.txt file as the target host",
	Long:  "SSH password free configuration, with the ip.txt file as the target host",
	RunE: func(cmd *cobra.Command, args []string) error {

		// 检查 pwd 参数是否为空
		if srcFile == "" || dstFile == "" {
			return fmt.Errorf("local path or target path required")
		}

		// 初始化ssh client对象，根据IP.txt
		cli := host.NewClientsConf()
		user := cli.GetUser()

		if pIPFile == "" {
			pIPFile = "ip.txt"
		}
		// 初始化
		cli.InitClientForKey(pIPFile, user, 22)

		// 初始化 sftp，并进行文件拷贝
		hosts := cli.HostList
		host.InitSFTP(hosts, srcFile, dstFile)

		return nil
	},
}

func init() {
	PushFile.PersistentFlags().StringVarP(&srcFile, "src", "s", "", "local file")
	PushFile.PersistentFlags().StringVarP(&dstFile, "dst", "d", "", "target file")
	PushFile.PersistentFlags().StringVarP(&pIPFile, "file", "f", "ip.txt", "ip list file")

	RootCmd.AddCommand(PushFile)
}