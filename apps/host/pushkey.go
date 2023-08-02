package host

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
)

// 推送密钥到远程服务器

type PushSSHKey struct {
	HostList []*ClientConfig
}

func NewPushSSHKey() *PushSSHKey {
	return &PushSSHKey{}
}

func (p *PushSSHKey) PushKey() {
	var (
		accHost []string
		errHost []string
	)

	publicKeyStr, err := p.getPrivateKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ins := range p.HostList {
		// 生成ssh客户端
		err := ins.CreateClient(context.Background())
		if err != nil {
			fmt.Printf("host %s: %s\n", ins.Host, err)
			errHost = append(errHost, ins.Host)
			continue
		}

		// 准备一个会话，执行shell命令
		sessionHome, err := ins.Client.NewSession()
		if err != nil {
			fmt.Println(err)
		}

		defer sessionHome.Close()
		// 首先获取远程主机用户的家目录
		homeCmd := "echo $HOME"
		homeDir, _ := sessionHome.Output(homeCmd)
		knownHostsDir := path.Join(strings.TrimRight(string(homeDir), "\r\n"), ".ssh")
		knownHostsPath := path.Join(knownHostsDir, "authorized_keys")
		cmd := fmt.Sprintf("mkdir -p %s && chmod 700 %s && echo %s >> %s && chmod 600 %s", knownHostsDir, knownHostsDir, publicKeyStr, knownHostsPath, knownHostsPath)

		// 新启动一个session，执行密钥签发。（由于这里的密钥配置需要依赖上面的home命令信息，所以分开写）
		session, err := ins.Client.NewSession()
		if err != nil {
			fmt.Println(err)
		}
		defer session.Close()
		_ = session.Run(cmd)

		accHost = append(accHost, ins.Host)
	}

	if len(accHost) > 0 {
		log.Printf("success: {num: %d, hosts: %s}\n", len(accHost), accHost)
		fmt.Printf("success: {num: %d, hosts: %s}\n", len(accHost), accHost)
	} else if len(errHost) > 0 {
		fmt.Printf("failed: {num: %d, hosts: %s}\n", len(errHost), errHost)
	}
}

// Init 读取IP文件，生成一个[]*ClientConfig 切片
func (p *PushSSHKey) Init(ipFile string, user, pwd string, port int64) {
	fileObj, err := os.Open(ipFile)
	if err != nil {
		fmt.Printf("open file failed, %s\n", err)
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		ins := &ClientConfig{
			Host:     "",
			Port:     port,
			UserName: user,
			Password: pwd,
		}
		if err != nil {
			if err == io.EOF {
				// 到达文件末尾，最后一行没有'\n'
				if line != "" {
					sLine := splitStr(line)
					if len(sLine) > 0 {
						ins.Host = sLine[0]
					}
					p.HostList = append(p.HostList, ins)
				}
			}
			return
		}
		sLine := splitStr(line)
		if len(sLine) > 0 {
			ins.Host = sLine[0]
		}
		p.HostList = append(p.HostList, ins)
	}
}

// splitStr 字符串切割 " 10.109.176.88 " --> [10.109.176.88]
func splitStr(str string) []string {
	return strings.Fields(str)
}

// getPrivateKey 获取当前用户的公钥
func (p *PushSSHKey) getPrivateKey() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get user home dir, %s", err)
	}

	publicKeyFile := path.Join(currentUser.HomeDir, ".ssh/id_rsa.pub")
	ok := p.isExistFile(publicKeyFile)
	if !ok {
		return "", fmt.Errorf("user %s did not generate an SSH key", currentUser.Name)
	}

	// 文件如果存在，读取文件中的内容
	publicKeyContent, err := os.ReadFile(publicKeyFile)
	if err != nil {
		return "", fmt.Errorf("read publickey file failed, %s", err)
	}

	publicKey := strings.TrimRight(string(publicKeyContent), "\r\n")
	return publicKey, nil
	//return string(publicKeyContent), nil
}

// IsExistFile 判断文件或目录是否存在，存在返回true，不存在返回false
func (p *PushSSHKey) isExistFile(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}