package main

import (
	"fmt"

	"github.com/Ibukun-tech/cloudcourier"
	file "github.com/Ibukun-tech/cloudcourier/File"
)

func main() {
	// _, err := os.Stat("../ffod/dh")
	// fmt.Println(err)
	genr, err := cloudcourier.NewCloudCourier(&cloudcourier.CloudCourierBridge{
		CloudProvider: cloudcourier.CloudinaryServices,
		ApiKey:        "Ibukunoluwa",
		ApiSecret:     "Oyetunji",
		CloudName:     "IbukunImage",
	})
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	v := genr.UploadFile(file.File{
		Path: "./Image/20171451.jpg",
		// FileName: "20171451.jpg",
	})
	fmt.Println(v)

	fmt.Println(genr)
}
