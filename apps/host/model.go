package host

import "golang.org/x/crypto/ssh"

type ClientConfig struct {
	Host     string
	Port     int64
	UserName string
	Password string
	Client   *ssh.Client
	Session  *ssh.Session
}