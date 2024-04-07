package cloudcourier

import (
	cld "github.com/Ibukun-tech/cloudcourier/cloudinary"
	"github.com/cloudinary/cloudinary-go/v2"
)

func cloudinaryFuncMiddleService(ccb *CloudCourierBridge) (General, error) {
	if ccb.ApiKey == "" && ccb.ApiSecret == "" && ccb.CloudName == "" {
		// Lets leave it like this for now still working on it and I want to modify to be possible and configured for what I want
		return nil, nil
	}

	client, err := cloudinary.NewFromParams(ccb.CloudName, ccb.ApiKey, ccb.ApiSecret)
	if err != nil {
		// Lets leave it like this for now still working on it and I want to modify to be possible and configured for what I want
		return nil, nil
	}
	// For now just return this
	return &cld.Cloudinary{
		Client: client,
	}, nil
}
