package main

import (
	"fmt"
	"log"

	"github.com/samuel-adekunle/cloudcourier"
)

func main() {
	client, err := cloudcourier.NewProviderClient(&cloudcourier.AWSProviderConfig{
		Bucket: "bucket-1",
		Region: "xx-north-1",
	})
	if err != nil {
		log.Fatalln(err)
	}

	files, err := client.ListFiles("example/directory")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Files:")
	for _, v := range files {
		fmt.Println(v)
	}
}
