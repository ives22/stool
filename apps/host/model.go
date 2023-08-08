package host

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"os"
)

type ClientConfig struct {
	Host       string
	Port       int64
	UserName   string
	Password   string
	PrivateKey ssh.Signer // 这里存放的是用于连接的私钥
	Client     *ssh.Client
	Session    *ssh.Session
	SFTPClient *sftp.Client // ftp的客户端
}

// ClientsConf 用于存放每个连接好的SSH Client
type ClientsConf struct {
	HostList []*ClientConfig
}

func NewClientsConf() *ClientsConf {
	return &ClientsConf{}
}

// FileInfo 本地文件信息， 用于文件下发时使用。
type FileInfo struct {
	Content *os.File    // 文件打开后的内容
	Model   os.FileMode // 文件的权限
	Size    int64       // 文件大小
}