package cloudcourier

import (
	"fmt"
	"io"
)

type Provider int

const (
	AWS Provider = iota
	GCS
	CLOUDINARY
)

type ProviderConfig interface {
	GetProvider() Provider
}

type ProviderClient interface {
	UploadFile(filePath string, reader io.Reader) error // Uploads a file from an io.Reader source to the specified path.
	DeleteFile(fileID string) error                     // Deletes a file identified by a unique identifier.
	ListFiles(directory string) ([]string, error)       // Lists files under a specified directory.
	GetFile(fileID string) (io.Reader, error)           // Retrieves a file as an io.Reader by its unique identifier.
}

func NewProviderClient(config ProviderConfig) (ProviderClient, error) {
	switch config.GetProvider() {
	case AWS:
		awsConfig, ok := config.(*AWSProviderConfig)
		if !ok {
			return nil, fmt.Errorf("Malformed AWS Provider Config: %v", config)
		}
		client, err := NewAWSProviderClient(awsConfig)
		if err != nil {
			return nil, err
		}
		return client, nil
	default:
		return nil, fmt.Errorf("Cloud provider in config is unsupported: %v", config)
	}
}
