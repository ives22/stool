package host_test

import (
	"context"
	"fmt"
	"github.com/ives22/stool/apps/host"

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
	p := host.NewPushSSHKey()
	p.Init("../../etc/ip.txt", "root", "ives", 22)
	p.PushKey()
	//hostList, err := p.Init("../../etc/ip.txt")
	//if should.NoError(err) {
	//	fmt.Println(hostList)
	//}
}