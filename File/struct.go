package file

type File struct {
	// This is the path of the file
	Path string `json:"path"`
	// This is the file name of the file to be stored
	FileName string `json:"fileName"`
}
