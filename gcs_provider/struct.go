package gcs_provider

import "cloud.google.com/go/storage"

type GcsClient struct {
	// The client we will use in communicating with the gcs
	Client *storage.Client
	// The name of the bucket to operate on
	BucketName string
}
