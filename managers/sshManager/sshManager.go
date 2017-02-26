package sshManager

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"fmt"
	"bytes"
)

func ExecuteCmd(command string, host string, port string, pvtKeyFilePath string, user string) (string, error) {
	session, err := connect(host, port, pvtKeyFilePath, user)
	if err != nil {
		return "", fmt.Errorf("Failed to create session: %s", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	err = session.Run(command)

	return b.String(), err
}

func connect(host string, port string, pvtKeyFilePath string, user string) (*ssh.Session, error) {

	var auths []ssh.AuthMethod
	auths = append(auths, publicKeyFile(pvtKeyFilePath))
	config := &ssh.ClientConfig{
		User: user,
		Auth: auths,
	}

	connection, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}

	return connection.NewSession()

}

func publicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}
