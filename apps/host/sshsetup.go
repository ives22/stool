package host

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

type SSHSetup struct {
	Home string
}

func NewSSHSetup() *SSHSetup {
	return &SSHSetup{}
}

// KeyGen 用于生成ssh密钥，需要指定用户和home
func (s *SSHSetup) KeyGen() {
	// ssh-keygen -t rsa -N '' -f ~/.ssh/id_rsa
	// 如果home传入的是 ~ 则通过echo $HOME 获取绝对路径
	if s.Home == "~" {
		// 这里采用os.ExpandEnv来对shell变量进行解释
		cmd := exec.Command("echo", []string{os.ExpandEnv("$HOME")}...)
		out, _ := cmd.Output()
		s.Home = strings.TrimSpace(string(out))
	}

	keyFile := path.Join(s.Home, ".ssh/id_rsa")

	// 要执行的命令和参数
	command := fmt.Sprintf("ssh-keygen -t rsa -f %s -N ''", keyFile)

	// 执行命令并捕获输出
	ret, err := s.ExecCmd(command)
	if err != nil {
		fmt.Printf("run cmd failed, %s", err)
	}
	fmt.Println(ret)
}

// ExecCmd 执行命令
func (s *SSHSetup) ExecCmd(command string) (string, error) {
	// 创建cmd对象
	//cmd := exec.Command(cmdName, cmdArgs...)
	//cmd := exec.Command("sh", "-c", "ssh-keygen -t rsa -f /home/test1/.ssh/id_rsa -N ''")
	cmd := exec.Command("sh", "-c", command)
	fmt.Println("cmd:", cmd)

	// 执行命令并捕获输出
	output, err := cmd.CombinedOutput()
	//cmd.Run()
	if err != nil {
		return "", err
	}
	return string(output), nil
}