package main

import (
	"fmt"

	"github.com/Ibukun-tech/cloudcourier"
)

func main() {
	s := &cloudcourier.CloudCourierBridge{
		CloudProvider: cloudcourier.CloudinaryServices,
		ApiKey:        "KSKSKSK",
		ApiSecret:     "KJJSJSJSJ",
		CloudName:     "Ibukun first Name",
	}

	t, _ := cloudcourier.NewCloudCourier(s)
	fmt.Println(t.UploadFile(&cloudcourier.File{}))
}
