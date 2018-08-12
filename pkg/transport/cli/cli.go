package cli

import (
	"golang.org/x/crypto/ssh/terminal"
	"fmt"
	"os"
)

type CLI struct {

}

func (c CLI) ShowCLI() {
	fd := int(os.Stdout.Fd())
	width, _, err := terminal.GetSize(fd)

	if err != nil {
		fmt.Printf(err.Error())
	}

	if width % 2 == 0 {
		width--
	}

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
}