package cloudcourier

import "github.com/Ibukun-tech/cloudcourier/config"

const (
	provider = map[string]string{
		CloudinaryServices: "cloudinary"
	}
	// This is for the services of the cloudinary
	CloudinaryServices = config.CloudinaryServices
)

var cloudCourierBridge = "CloudCourierBridge"
