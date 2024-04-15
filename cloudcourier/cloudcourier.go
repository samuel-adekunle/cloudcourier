package cloudcourier

import (
	"fmt"
	"io"
)

// CloudCourierBridge represents the connection settings for a cloud storage provider.
//
// It is designed to accommodate various cloud services by storing essential credentials.
type CloudCourierBridge struct {
	CloudRegion   string // used by AWS S3 to specify the Region
	CloudProvider string // Identifier for the cloud service provider
	ApiKey        string // API key for services that require it, such as Cloudinary
	ApiSecret     string // API secret corresponding to the ApiKey for added security
	CloudName     string // Specific to Cloudinary; this is the name of the cloud to connect to
	CloudBucket   string // Used by AWS S3, Google Cloud Storage, etc., to specify the storage bucket
}

// StorageClient interface defines the behavior for cloud storage operations.
//
// This abstraction allows for the implementation of various cloud services under a unified interface.
type StorageClient interface {
	UploadFile(filePath string, reader io.Reader) error // Uploads a file from an io.Reader source to the specified path.
	DeleteFile(fileID string) error                     // Deletes a file identified by a unique identifier.
	ListFiles(directory string) ([]string, error)       // Lists files under a specified directory.
	GetFile(fileID string) (io.Reader, error)           // Retrieves a file as an io.Reader by its unique identifier.
}

// NewCloudCourier initializes a client for the specified cloud service provider.
// It validates necessary parameters and returns a client adhering to the StorageClient interface.
func NewCloudCourier(cbb *CloudCourierBridge) (StorageClient, error) {
	if cbb == nil {
		return nil, fmt.Errorf("cloud courier configuration cannot be nil")
	}

	switch cbb.CloudProvider {
	case "cloudinary":
		return newCloudinaryClient(cbb)
	case "aws":
		return newAWSClient(cbb)
	case "gcs":
		return newGcsClient(cbb)
	default:
		return nil, fmt.Errorf("unsupported cloud provider: %s", cbb.CloudProvider)
	}
}
