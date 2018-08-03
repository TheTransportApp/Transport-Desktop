package sftp

import "fmt"

type SftpCommand struct {
	Host string
	Port int
	Username string
	Password string

	remote Remote
}

func (cmd SftpCommand) Connect() {
	remote := Remote{}

	result := remote.Connect(cmd.Host, cmd.Port, cmd.Username, cmd.Password)

	if result {
		fmt.Println("Connection established")
	} else {
		fmt.Println("Connection failed")
	}
}