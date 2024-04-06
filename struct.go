package cloudcourier

import "github.com/Ibukun-tech/cloudcourier/types"

type CloudCourierBridge = types.CloudCourierBridge

type General interface {
	UploadFile(file interface{}) error
}
