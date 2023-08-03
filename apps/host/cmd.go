package host

import (
	"bufio"
	"context"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"os/user"
	"path"
)

// Command 执行命令，依赖于先配置免密后才行
func (c *ClientsConf) Command(cmd string) {
	for _, ins := range c.HostList {
		err := ins.CreateClientForSecretKey(context.Background())
		if err != nil {
			fmt.Printf("host %s: %s\n", ins.Host, err)
			continue
		}
		defer ins.Client.Close()

		session, err := ins.Client.NewSession()
		if err != nil {
			fmt.Printf("host %s: %s\n", ins.Host, err)
			continue
		}
		defer session.Close()

		cmdOutput, err := session.CombinedOutput(cmd)
		if err != nil {
			fmt.Printf("host %s, %s\n", ins.Host, err)
			continue
		}
		fmt.Printf("==================== %s ====================\n", ins.Host)
		fmt.Println(string(cmdOutput))
	}
}

// InitClientForKey 初始化每个client，根据获取当前用户的公钥后，然后获取传递给每个client对象上，用于建立免密连接
func (c *ClientsConf) InitClientForKey(ipFile, user string, port int64) {
	// 获取当前用户的私钥
	signer, err := c.getPrivateKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileObj, err := os.Open(ipFile)
	if err != nil {
		fmt.Printf("open file failed, %s\n", err)
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		ins := &ClientConfig{
			Port:       port,
			UserName:   user,
			PrivateKey: signer,
		}
		if err != nil {
			if err == io.EOF {
				if line != "" {
					sLine := splitStr(line)
					if len(sLine) > 0 {
						ins.Host = sLine[0]
					}
					c.HostList = append(c.HostList, ins)
				}
			}
			return
		}
		sLine := splitStr(line)
		if len(sLine) > 0 {
			ins.Host = sLine[0]
		}
		c.HostList = append(c.HostList, ins)
	}
}

// getPrivateKey 获取当前用户的私钥文件路径
func (c *ClientsConf) getPrivateKey() (ssh.Signer, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home dir, %s", err)
	}

	privateKeyFile := path.Join(currentUser.HomeDir, ".ssh/id_rsa")
	ok := c.isExistFile(privateKeyFile)
	if !ok {
		return nil, fmt.Errorf("user %s did not generate an SSH key", currentUser.Name)
	}

	// 文件如果存在，读取文件中的内容
	privateKeyContent, err := os.ReadFile(privateKeyFile)
	//fmt.Println(string(privateKeyContent))
	if err != nil {
		return nil, fmt.Errorf("read privateKey file failed, %s", err)
	}

	signer, err := ssh.ParsePrivateKey(privateKeyContent)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key, %s", err)
	}
	return signer, nil
}