package sftp

import (
	"fmt"
	commandline "github.com/thetransportapp/transport-desktop/pkg/transport/cli"
)

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

	cli := commandline.CLI{}
	cli.ShowCLI()
}