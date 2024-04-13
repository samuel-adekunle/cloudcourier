package cloudcourier

import (
	"fmt"
	"github.com/Ibukun-tech/cloudcourier/aws_provider"
)

func newAWSClient(cbb *CloudCourierBridge) (StorageClient, error) {
	if cbb.ApiKey == "" || cbb.ApiSecret == "" || cbb.CloudBucket == "" {
		return nil, fmt.Errorf("incomplete AWS configuration")
	}
	// Assuming AWS SDK setup here
	return &aws_provider.AWSClient{ /* configured client */ }, nil
}
