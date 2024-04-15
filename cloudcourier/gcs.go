package cloudcourier

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/Ibukun-tech/cloudcourier/gcs_provider"
)

func newGcsClient(ccb *CloudCourierBridge) (StorageClient, error) {
	if ccb.CloudBucket == "" {
		return nil, fmt.Errorf("no bucket name for google cloud storage")
	}
	var ctx context.Context
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return &gcs_provider.GcsClient{
		Client:     client,
		BucketName: ccb.CloudBucket,
	}, nil
}
