package gcs_provider

import "io"

// 	UploadFile(filePath string, reader io.Reader) error // Uploads a file from an io.Reader source to the specified path.
// DeleteFile(fileID string) error                     // Deletes a file identified by a unique identifier.
// ListFiles(directory string) ([]string, error)       // Lists files under a specified directory.
// GetFile(fileID string) (io.Reader, error)           // Retrieves a file as an io.Reader by its unique identifier.

func (g *GcsClient) UploadFile(filepath string, reader io.Reader) error {
	return nil
}

func (g *GcsClient) DeleteFile(fieldID string) error {
	return nil
}
func (g *GcsClient) ListFiles(directory string) ([]string, error) {
	var st []string
	st = append(st, "jsj")
	return st, nil
}

func (g *GcsClient) GetFile(fileID string) (io.Reader, error) {
	return nil, nil
}
