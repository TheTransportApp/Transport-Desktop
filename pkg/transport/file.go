package transport

type FileInfo struct {
	path string
	name string
	size int64
}

func (f FileInfo) GetPath() string {
	return f.path
}

func (f FileInfo) GetName() string {
	return f.name
}

func (f FileInfo) GetSize() int64 {
	return f.size
}