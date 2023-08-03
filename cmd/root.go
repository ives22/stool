package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var (
	verbose    bool
	initIPFile bool
)

func GenerateIPFile() {
	// 获取当前的路径
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	filePath := path.Join(dir, "ip.txt")
	fileObj, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()

	fileObj.WriteString("192.168.1.2\n192.168.1.3\n192.168.1.4\n")
	fmt.Printf("create ip.txt success, path: %s\n", filePath)
}

var RootCmd = &cobra.Command{
	Use:               "stool",
	Short:             "stool: 运维小工具，用于ssh密钥生成、免密配置、shell命令执行",
	Long:              "stool: 运维小工具，用于ssh密钥生成、免密配置、shell命令执行",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Run: func(cmd *cobra.Command, args []string) {
		// 在Run函数中根据-v参数的值输出详细信息
		if verbose {
			fmt.Println("stool version: v1.0.0")
			return
		}

		if initIPFile {
			GenerateIPFile()
			return
		}
		_ = cmd.Help()
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "version", "v", false, "show stool version")
	RootCmd.PersistentFlags().BoolVarP(&initIPFile, "init", "i", false, "Generate ip.txt template files")
}