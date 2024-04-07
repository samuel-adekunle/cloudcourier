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

	if ccb.CloudProvider == CloudinaryServices {
		continue
	} else {
	}
	return &cld.Cloudinary{}, nil
}
