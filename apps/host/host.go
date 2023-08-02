package host

import (
	"context"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func NewClientConfig(host string, port int64, user, pass string) *ClientConfig {
	return &ClientConfig{
		Host:     host,
		Port:     port,
		UserName: user,
		Password: pass,
	}
}

// CreateClient 创建ssh client
func (c *ClientConfig) CreateClient(ctx context.Context) error {
	config := &ssh.ClientConfig{
		User: c.UserName,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return fmt.Errorf("failed to dial, %v", err)
	}

	c.Client = client
	return nil
}

func (c *ClientConfig) RunShell(ctx context.Context, cmd string) (ret string, err error) {
	// 准备一个会话，执行shell命令
	session, err := c.Client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed create session, %v", err)
	}
	defer session.Close()

	// 执行shell命令
	output, err := session.CombinedOutput(cmd)

	if err != nil {
		log.Fatalf("无法执行命令: %v", err)
	}

	fmt.Println(string(output))
	return string(output), nil
}

func (c *ClientConfig) RunShellOne(cmd string) (ret string, err error) {
	// 准备一个会话，执行shell命令
	session, err := c.Client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed create session, %v", err)
	}
	defer session.Close()

	// 执行shell命令
	output, err := session.CombinedOutput(cmd)

	if err != nil {
		log.Fatalf("无法执行命令: %v", err)
	}

	fmt.Println(string(output))
	return string(output), nil
}