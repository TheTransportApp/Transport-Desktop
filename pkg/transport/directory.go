package transport

type DirectoryInfo struct {
	path string
	name string
	files []FileInfo
	directories []DirectoryInfo
}

func (d DirectoryInfo) GetPath() string {
	return d.path
}

func (d DirectoryInfo) GetName() string {
	return d.name
}

func (d DirectoryInfo) GetFiles() []FileInfo {
	return d.files
}

func (d DirectoryInfo) GetDirectories() []DirectoryInfo {
	return d.directories
}
