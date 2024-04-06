package cloudcourier

import cld "github.com/Ibukun-tech/cloudcourier/cloudinary"

func NewCloudCourier(ccb CloudCourierBridge) (General, error) {
	return &cld.Cloudinary{}, nil
}
