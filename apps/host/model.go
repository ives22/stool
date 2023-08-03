package host

import "golang.org/x/crypto/ssh"

type ClientConfig struct {
	Host       string
	Port       int64
	UserName   string
	Password   string
	PrivateKey ssh.Signer // 这里存放的是用于连接的私钥
	Client     *ssh.Client
	Session    *ssh.Session
}

type ClientsConf struct {
	HostList []*ClientConfig
}

func NewClientsConf() *ClientsConf {
	return &ClientsConf{}
}