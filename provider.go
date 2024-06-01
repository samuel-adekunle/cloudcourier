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

type ProviderConstructor func(ProviderConfig) (ProviderClient, error)

var providerConstructors map[Provider]ProviderConstructor

func RegisterProviderConstructor(provider Provider, constructor ProviderConstructor) {
	if providerConstructors == nil {
		providerConstructors = make(map[Provider]ProviderConstructor)
	}
	providerConstructors[provider] = constructor
}

func NewProviderClient(config ProviderConfig) (ProviderClient, error) {
	constructor, supported := providerConstructors[config.GetProvider()]
	if !supported {
		return nil, fmt.Errorf("cloud provider in config is unsupported: %v", config)
	}
	client, err := constructor(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create %v client: %v", config.GetProvider(), err)
	}
	return client, nil
}
