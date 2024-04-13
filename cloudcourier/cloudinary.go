package cloudcourier

import (
	"errors"
	cloudinaryclient "github.com/Ibukun-tech/cloudcourier/cloudinary"
)

func newCloudinaryClient(cbb *CloudCourierBridge) (StorageClient, error) {
	if cbb.ApiKey == "" || cbb.ApiSecret == "" || cbb.CloudName == "" {
		return nil, errors.New("incomplete Cloudinary configuration") // TODO: handle properly.
	}

	return &cloudinaryclient.CloudinaryClient{ /* configured client */ }, nil
}
