package gcs_provider

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
)

//	UploadFile(filePath string, reader io.Reader) error // Uploads a file from an io.Reader source to the specified path.
//
// DeleteFile(fileID string) error                     // Deletes a file identified by a unique identifier.
// ListFiles(directory string) ([]string, error)       // Lists files under a specified directory.
// GetFile(fileID string) (io.Reader, error)           // Retrieves a file as an io.Reader by its unique identifier.

func (g *GcsClient) UploadFile(filePath string, reader io.Reader) error {
	var BaseFileName string
	ctx := context.Background()
	if filePath != "" {
		BaseFileName = filepath.Base(filePath)
	}
	if filePath == "" {
		return fmt.Errorf("you did not specify the filepath")
	}
	obj := g.Client.Bucket(g.BucketName).Object(BaseFileName)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, reader); err != nil {
		return fmt.Errorf("you did not set the reader to the file")
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("could not upload file")
	}
	v, err := obj.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println(v)
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
