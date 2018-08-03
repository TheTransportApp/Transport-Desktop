package transport

import "os"

// Remote defines a type of remote host for use in the client
type Remote interface {
	// Get the name of the transport type
	GetName() string

	// Connect to server
	Connect(host string, port int, username string, password string) bool

	// List the directory structure of the server
	ListDirectoryStructure()

	// List the content of the current directory
	ListContent()

	// Copy a local file to remote server
	CopyToRemote(source string, destination string) bool

	// Copy a remote file to local path
	CopyToLocal(source string, destination string) bool

	// Rename file on remote server
	Rename(oldName string, newName string) bool

	// Delete file on remote server
	Delete(path string) bool

	// Create file on remote server
	CreateFile(path string) bool

	// Create folder on remote server
	CreateFolder(path string) bool
}
