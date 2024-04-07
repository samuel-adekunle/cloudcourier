package cloudcourier

import (
	file "github.com/Ibukun-tech/cloudcourier/File"
	"github.com/Ibukun-tech/cloudcourier/types"
)

type CloudCourierBridge types.CloudCourierBridge

type File file.File

type General interface {
	UploadFile(file interface{}) error
}
