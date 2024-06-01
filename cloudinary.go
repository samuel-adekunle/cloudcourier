package cloudcourier

import (
	"fmt"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
)

type CloudinaryProviderConfig struct {
	Tag       string
	ApiKey    string
	ApiSecret string
	CloudName string
}

func (c *CloudinaryProviderConfig) GetProvider() Provider {
	return AWS
}

type CloudinaryProviderClient struct {
	Config *CloudinaryProviderConfig
	Client *cloudinary.Cloudinary
}

func newCloudinaryProviderClient(config ProviderConfig) (ProviderClient, error) {
	cloudinaryConfig, ok := config.(*CloudinaryProviderConfig)
	if !ok {
		return nil, fmt.Errorf("incorrect provider config provided")
	}
	if cloudinaryConfig.ApiKey == "" || cloudinaryConfig.ApiSecret == "" || cloudinaryConfig.CloudName == "" {
		return nil, fmt.Errorf("you did not specify either the api key, the api secret and also the cloud name")
	}
	client, err := cloudinary.NewFromParams(cloudinaryConfig.CloudName, cloudinaryConfig.ApiKey, cloudinaryConfig.ApiSecret)
	if err != nil {
		return nil, err
	}
	return &CloudinaryProviderClient{
		Config: cloudinaryConfig,
		Client: client,
	}, nil
}

func init() {
	RegisterProviderConstructor(CLOUDINARY, newCloudinaryProviderClient)
}

func (c *CloudinaryProviderClient) UploadFile(filePath string, reader io.Reader) error {
	return fmt.Errorf("Unimplemented")
}

func (c *CloudinaryProviderClient) DeleteFile(fileID string) error {
	return fmt.Errorf("Unimplemented")
}

func (c *CloudinaryProviderClient) ListFiles(directory string) ([]string, error) {
	return nil, fmt.Errorf("Unimplemented")
}

func (c *CloudinaryProviderClient) GetFile(fileID string) (io.Reader, error) {
	return nil, fmt.Errorf("Unimplemented")
}
