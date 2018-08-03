package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thetransportapp/transport-desktop/pkg/transport/sftp"
	"bufio"
	"os"
	"strconv"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

var host string
var port int
var username string

var sftpCmd = &cobra.Command{
	Use:   "sftp",
	Short: "SFTP Connection",
	Long: `With this command you can build a connection to a SFTP server`,
	Run: func(cmd *cobra.Command, args []string) {
		if host == "" {
			host = readFromConsole("Host")
		}

		if port == 0 {
			if p, err := strconv.Atoi(readFromConsole("Port")); err == nil {
				port = p
			} else {
				fmt.Println(err)
				return
			}
		}

		if username == "" {
			username = readFromConsole("Username")
		}

		password := readPassword()

		fmt.Println()

		command := sftp.SftpCommand{
			Host: host,
			Port: port,
			Username: username,
			Password: password,
		}
		command.Connect()
	},
}

func init() {
	RootCmd.AddCommand(sftpCmd)

	sftpCmd.Flags().StringVar(&host, "host", "", "IP address of host")
	sftpCmd.Flags().IntVarP(&port, "port", "p", 0, "Port of host")
	sftpCmd.Flags().StringVarP(&username, "username", "u", "", "Username to login")
}

func readFromConsole(name string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(name + ": ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	return text
}

func readPassword() string {
	fmt.Print("Password: ")
	pw, err := terminal.ReadPassword(int(syscall.Stdin))

	if err == nil {
		return string(pw)
	}
	return ""
}