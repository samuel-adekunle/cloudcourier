package file_handling

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

// File represents a single file to be uploaded.
type File struct {
	ToHandle io.Reader // Handle to red the file data.
	Path     string    // Path of the file on the local disk.
	FileName string    // Name of the file to be stored in the cloud.
	FileSize int64     // Size of the file in bytes
}

// NewFile creates a new File instance from a path.
//
// It validates the path, prepares the file handle, and checks file size.
func NewFile(path string) (*File, error) {
	if path == "" {
		return nil, errors.New("file path cannot be empty")
	}

	info, err := os.Stat(path)
	if err != nil {
		return nil, err // Path does not exist or other error. TODO: handle the error properly.
	}

	if info.IsDir() {
		return nil, errors.New("specified path is a directory, not a file")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err // Error opening file. TODO: handle error properly.
	}

	// Ensure file is closed if there are subsequent errors
	defer func() {
		if err != nil {
			file.Close()
		}
	}()

	return &File{
		ToHandle: file,
		Path:     path,
		FileName: filepath.Base(path),
		FileSize: info.Size(),
	}, nil
}

// CheckFileSize checks if the File exceeds a given size limit in bytes.
//
// This is useful because some Cloud Providers gate the size of files that
// come onto them, and also to prevent you from getting excessive billing at
// the end of the month, because your users uploaded large files.
func (f *File) CheckFileSize(limit int64) error {
	if f.FileSize > limit {
		return errors.New("file size exceeds the maximum allowed limit")
	}

	return nil
}

// Use a proper logging library for handling the errors.
// Add your other functions. Usage is like this:
/*
func main() {
	filePath := "path/to/your/file.txt"
	file, err := file_handling.NewFile(filePath)
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}

	// Check file size against a limit (e.g., 10 MB)
	sizeLimit := int64(10 * 1024 * 1024) // 10 MB in bytes
	if err := file.CheckFileSize(sizeLimit); err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Println("File is within size limits and ready for upload.")
}
*/
