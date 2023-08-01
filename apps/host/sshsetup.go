package host

import (
	"fmt"
	"os/exec"
)

type SSHSetup struct {
	Home string
}

func NewSSHSetup() *SSHSetup {
	return &SSHSetup{}
}

// KeyGen 用于生成ssh密钥，需要指定用户和home
func (s *SSHSetup) KeyGen() {
	// ssh-keygen -t rsa -f ~/.ssh/id_rsa
	keyFile := fmt.Sprintf("%s/.ssh/id_rsa", s.Home)

	// 要执行的命令和参数
	cmdName := "ssh-keygen"
	cmdArgs := []string{"-t", "rsa", "-f", keyFile}

	// 执行命令并捕获输出
	ret, err := s.ExecCmd(cmdName, cmdArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)

}

// ExecCmd 执行命令
func (s *SSHSetup) ExecCmd(cmdName string, cmdArgs []string) (string, error) {
	// 创建cmd对象
	cmd := exec.Command(cmdName, cmdArgs...)
	// 执行命令并捕获输出
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func main() {
	//cmd := exec.Command("ls", "/root")
	//output, err := cmd.Output()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(output))

	c := NewSSHSetup()
	c.Home = "/home/test1"
	c.KeyGen()
}