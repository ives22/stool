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
	cli := host.NewClientConfig("127.0.0.1", 22, "root", "")
	err := cli.CreateClient(context.Background())
	fmt.Println("haha")
	if should.NoError(err) {
		cli.RunShell(context.Background(), "ls /root")
	}
}