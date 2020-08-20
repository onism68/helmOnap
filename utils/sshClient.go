package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"time"
)

type SSH struct {
	Host, User, Password               string
	IdentityFile, IdentityFilePassword string
	Port                               int
}

func NewSSH(host, user, password, identityFile, idfilePass string, port int) *SSH {
	return &SSH{
		Host:                 host,
		User:                 user,
		Password:             password,
		Port:                 port,
		IdentityFile:         identityFile,
		IdentityFilePassword: idfilePass,
	}
}

func (s *SSH) Connect() (*ssh.Session, error) {
	auths := []ssh.AuthMethod{
		ssh.Password(s.Password),
	}
	key, err := ioutil.ReadFile(s.IdentityFile)
	if err == nil {
		key, err := DecryptKey(key, []byte(s.IdentityFilePassword))
		if err != nil {
			goto M
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err == nil {
			auths = append(auths, ssh.PublicKeys(signer))
		}
	}
M:
	config := &ssh.ClientConfig{
		User:    s.User,
		Auth:    auths,
		Timeout: time.Second * 10,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}

	return client.NewSession()
}
