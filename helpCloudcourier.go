package cloudcourier

import (
	"fmt"

	cld "github.com/Ibukun-tech/cloudcourier/cloudinary"
	"github.com/cloudinary/cloudinary-go/v2"
)

func cloudinaryFuncMiddleService(ccb *CloudCourierBridge) (General, error) {
	var client *cloudinary.Cloudinary
	if ccb.ApiKey == "" && ccb.ApiSecret == "" && ccb.CloudName == "" {
		return nil, fmt.Errorf("you did not specify either the api key, the api secret and also the cloud name")
	}

	client, err := cloudinary.NewFromParams(ccb.CloudName, ccb.ApiKey, ccb.ApiSecret)
	if err != nil {
		return nil, err
	}
	// For now just return this
	return &cld.Cloudinary{
		CloudName: ccb.CloudName,
		Client:    client,
	}, nil
}
