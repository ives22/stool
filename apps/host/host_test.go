package host_test

import (
	"context"
	"fmt"
	"github.com/ives22/stool/apps/host"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"time"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunShell(t *testing.T) {
	should := assert.New(t)
	cli := host.NewClientConfig("", 22, "user1", "111")
	err := cli.CreateClient(context.Background())
	if should.NoError(err) {
		cli.RunShell(context.Background(), " ls -l /")
	}
}

func TestIsExistFile(t *testing.T) {
	s := host.NewSSHSetup()
	ok := s.IsExistFile("/Users/liyanjie/Documents/08 go/03 Project/stool")
	fmt.Println(ok)
}

func TestPushKey(t *testing.T) {
	//should := assert.New(t)
	p := host.NewClientsConf()
	p.Init("../../etc/ip.txt", "root", "ives", 22)
	p.PushKey()
	//hostList, err := p.Init("../../etc/ip.txt")
	//if should.NoError(err) {
	//	fmt.Println(hostList)
	//}
}

// TestScp 测试scp
func TestScp(t *testing.T) {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("ives.123"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	addr := fmt.Sprintf("%s:%d", "124.71.33.240", 22)
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conn)

	// 建立sftp客户端
	client, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println(err)
	}

	//clientFile, err := client.Create("/tmp/1.txt")
	info, err := client.Stat("/tmp222/1.txt")
	if os.IsNotExist(err) {
		fmt.Println("不存在")
	}
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("f", info)

	clientFile, err := client.Create("/tmp22/protoc-23.4-osx-aarch_64.zip")
	if err != nil {
		fmt.Println("打开远程文件失败", err)
	}
	defer clientFile.Close()

	//localFile, err := os.Open("/Users/liyanjie/Desktop/test.sh")
	localFile, err := os.Open("/Users/liyanjie/Desktop/protoc-23.4-osx-aarch_64.zip")
	if err != nil {
		fmt.Println("打开本地文件失败,", err)
	}
	defer localFile.Close()

	_, err = io.Copy(clientFile, localFile)
	if err != nil {
		fmt.Println("copy 文件失败")
	}

	d, f := sftp.Split("/tmp22/protoc-23.4-osx-aarch_64.zip")
	fmt.Println(d)
	fmt.Println(f)

}

func openLocalFile(path string) {
	localFile, err := os.Open(path)
	if err != nil {
		fmt.Println("打开本地文件失败,", err)
	}
	defer localFile.Close()
}