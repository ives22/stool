package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var verbose bool

var RootCmd = &cobra.Command{
	Use:               "stool",
	Short:             "stool: shell小工具，用于ssh密钥生成、免密配置、shell命令执行",
	Long:              "stool: shell小工具，用于ssh密钥生成、免密配置、shell命令执行",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Run: func(cmd *cobra.Command, args []string) {
		// 在Run函数中根据-v参数的值输出详细信息
		if verbose {
			fmt.Println("stool version: v1.0.0")
			return
		}
		_ = cmd.Help()
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "version", "v", false, "show stool version")
}