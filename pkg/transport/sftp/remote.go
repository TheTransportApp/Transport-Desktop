package sftp

import (
	"golang.org/x/crypto/ssh"
	"github.com/pkg/sftp"
	"fmt"
	"os"
	"io"
)

type Remote struct {
	conn *ssh.Client
	secureFtp *sftp.Client
}

func (r Remote) GetName() string {
	return "SFTP"
}

func (r Remote) Connect(host string, port int, username string, password string) bool {
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},

	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	var err error
	r.conn, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), sshConfig)

	if err != nil {
		return false
	}

	r.secureFtp, err = sftp.NewClient(r.conn)

	return err == nil
}

func (r Remote) ListDirectoryStructure() {
	walker := r.secureFtp.Walk(".")

	for walker.Step() {
		if walker.Err() != nil {
			continue
		}

	}
}

func (r Remote) ListContent() {

}

func (r Remote) CopyToRemote(source string, destination string) bool {
	file, err := r.secureFtp.Create(destination)

	if err != nil {
		return false
	}

	sourceFile, err := os.Open(source)

	if err != nil {
		return false
	}

	_, err = io.Copy(file, sourceFile)

	if err != nil {
		return false
	}

	sourceFile.Close()
	file.Close()

	return true
}

func (r Remote) CopyToLocal(source string, destination string) bool {
	file, err := r.secureFtp.Open(source)

	if err != nil {
		return false
	}

	destinationFile, err := os.Open(destination)

	if err != nil {
		return false
	}

	_, err = io.Copy(destinationFile, file)

	if err != nil {
		return false
	}

	destinationFile.Close()
	file.Close()

	return true
}

func (r Remote) Rename(oldName string, newName string) bool {
	err := r.secureFtp.Rename(oldName, newName)

	return err == nil
}

func (r Remote) Delete(path string) bool {
	fileInfo, err := r.secureFtp.Stat(path)

	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		err = r.secureFtp.RemoveDirectory(path)
	} else {
		err = r.secureFtp.Remove(path)
	}

	return err != nil
}

func (r Remote) CreateFile(path string) bool {
	_, err := r.secureFtp.Create(path)

	return err != nil
}

func (r Remote) CreateFolder(path string) bool {
	err := r.secureFtp.MkdirAll(path)

	return err != nil
}