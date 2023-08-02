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

// KeyGen 用于生成ssh密钥。ssh-keygen -t rsa -N ” -f ~/.ssh/id_rsa
func (s *SSHSetup) KeyGen() {

	var keyFile string

	// 如果home传入的是 ~ 则通过echo $HOME 获取绝对路径
	if s.Home == "~" {
		// 这里采用os.ExpandEnv来对shell变量进行解释
		cmd := exec.Command("echo", []string{os.ExpandEnv("$HOME")}...)
		out, _ := cmd.Output()
		s.Home = strings.TrimSpace(string(out))
		keyFile = path.Join(s.Home, ".ssh/id_rsa")
	} else {
		// 如果传入的是其它目录，比如： /tmp/abc，则判断该目录是否存在, 如果不存在则退出，这里不自动帮忙创建目录
		ok := s.IsExistFile(s.Home)
		if !ok {
			fmt.Printf("%s: no such file or directory\n", s.Home)
			return
		}
		keyFile = path.Join(s.Home, "id_rsa")
	}

	// 判断要生成的文件是否存在，如果存在，则提示是否继续要创建
	var isTrue string
	ok := s.IsExistFile(keyFile)
	if ok {
		fmt.Printf("%s already exists.\n", keyFile)
		fmt.Print("Overwrite (y/n)? ")
		fmt.Scanln(&isTrue)
		switch isTrue {
		case "y", "Y", "yes":
			_ = os.Remove(keyFile)
		case "n", "N", "no":
			return
		}
	}

	// 要执行的命令和参数
	command := fmt.Sprintf("ssh-keygen -t rsa -f %s -N ''", keyFile)

	// 执行命令并捕获输出
	ret, err := s.ExecCmd(command)
	if err != nil {
		fmt.Printf("run cmd failed, %s\n", err)
		return
	}

	fmt.Println(ret)
	fmt.Println("Generating success.")
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
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// IsExistFile 判断文件或目录是否存在，存在返回true，不存在返回false
func (s *SSHSetup) IsExistFile(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}