package domain

type FileInfo struct {
	ID   string
	Name string
	Size int64

	Location string
}

func NewFileInfo(ID string, name string, size int64, location string) FileInfo {
	return FileInfo{ID: ID, Name: name, Size: size, Location: location}
}
