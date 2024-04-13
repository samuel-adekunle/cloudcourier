package cloudinary_client

import (
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
)

// TODO: model this properly, implement the functions properly too.

type CloudinaryClient struct {
	// Name you give to a file to be stored in cloudinary
	Tag string
	// This is for the transport side responsible for the intercommunication
	Client *cloudinary.Cloudinary
	// This helps to show the cloud provider you want to use
	Provider string
	// The Api key for now I am just using it to access the cloudinary services
	ApiKey string
	// The Api secret which corresponds to the ApiKey
	ApiSecret string
	// This is for cloudinary you need to provide the cloud name for the cloudinary
	CloudName string
}

func (client CloudinaryClient) UploadFile(filePath string, reader io.Reader) error {
	//TODO implement me
	panic("implement me")
}

func (client CloudinaryClient) DeleteFile(fileID string) error {
	//TODO implement me
	panic("implement me")
}

func (client CloudinaryClient) ListFiles(directory string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (client CloudinaryClient) GetFile(fileID string) (io.Reader, error) {
	//TODO implement me
	panic("implement me")
}
