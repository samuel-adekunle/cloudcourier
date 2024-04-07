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
	if v.Elem().Name() != "" {
	}
	return &cld.Cloudinary{}, nil
}
