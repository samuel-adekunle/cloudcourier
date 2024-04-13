package cloudcourier

import (
	"fmt"

	"github.com/Ibukun-tech/cloudcourier/aws_provider"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func newAWSClient(cbb *CloudCourierBridge) (StorageClient, error) {
	if cbb.CloudBucket == "" || cbb.CloudRegion == "" {
		return nil, fmt.Errorf("incomplete AWS configuration")
	}
	sess, err := session.NewSession(&aws.Config{
		Region: &cbb.CloudRegion,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %v", err)
	}
	s3Svc := s3.New(sess)
	return &aws_provider.AWSClient{S3: s3Svc,
		BucketName: cbb.CloudBucket}, nil
}
