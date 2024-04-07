package cloudcourier

import (
	"fmt"
	"reflect"

	cld "github.com/Ibukun-tech/cloudcourier/cloudinary"
)

func NewCloudCourier(ccb *CloudCourierBridge) (General, error) {
	if ccb == nil {
		return nil, fmt.Errorf("it's an error")
	}

	v := reflect.TypeOf(ccb)
	if v.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("this is an inavalid type %s", v.Elem().Name())
	}

	if v.Elem().Name() != cloudCourierBridge {
		return nil, fmt.Errorf("invalid type the type to be use is %s", v.Elem().Name())
	}

	if ccb.CloudProvider == "" {
		return nil, fmt.Errorf("we need to ")
	}

	if _, ok := provider[ccb.CloudProvider]; !ok {
		return nil, fmt.Errorf("does not support that cloud services %s", ccb.CloudProvider)
	}

	// Other services require a bucket but cloudinary services do not require a bucket
	if ccb.CloudProvider == CloudinaryServices && ccb.CloudName == "" {
		return nil, fmt.Errorf("since we are using cloudinary service you need to specify the cloudname")
	}

	if ccb.CloudProvider != CloudinaryServices && ccb.CloudBucket == "" {
		// For now this is what I am thinking I will do but still working on how I will continue about it
		return nil, nil
	}
	switch ccb.CloudProvider {
	case CloudinaryServices:
		return cloudinaryFuncMiddleService(ccb)
	}
	return &cld.Cloudinary{}, nil
}
