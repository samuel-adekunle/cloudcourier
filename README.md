# Cloud Courier

Cloud Courier is a Go library designed to simplify the process of uploading files to various cloud storage services.
It provides a unified interface to interact with different cloud providers, allowing developers to integrate cloud storage capabilities into their applications seamlessly.

## Features

- **Unified API**: Use a single, consistent API to interact with multiple cloud storage providers.
- **Extensible**: Easily add support for additional cloud providers.
- **Simple Configuration**: Configure your cloud storage credentials and settings in one place.
- **Error Handling**: Robust error handling for common edge cases in file uploads.

## Supported Cloud Providers

- AWS
- Google Cloud Storage (coming soon)
- Cloudinary (coming soon)

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

To install Cloud Courier, use the following `go get` command:

```sh
go get -u github.com/samuel-adekunle/cloudcourier
```

### Usage

Here's a quick example of how to use Cloud Courier to list files in an AWS directory:

```go
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
```

See more examples in the [examples directory](/examples/)

## Acknowledgments

- Thanks to all the contributors who have helped with the development of Cloud Courier.

## Contact

For questions and feedback, please reach out to @samuel-adekunle.