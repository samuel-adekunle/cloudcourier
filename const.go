package cloudcourier

import "github.com/Ibukun-tech/cloudcourier/config"

var (
	provider = map[string]string{
		CloudinaryServices: "cloudinary",
	}
)

const (

	// This is for the services of the cloudinary
	CloudinaryServices = config.CloudinaryServices
)

var cloudCourierBridge = "CloudCourierBridge"
