package cloudcourier

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSProviderConfig struct {
	Bucket string
	Region string
}

func (c *AWSProviderConfig) GetProvider() Provider {
	return AWS
}

type AWSProviderClient struct {
	Config *AWSProviderConfig
	Client *s3.S3
}

func newAWSProviderClient(config ProviderConfig) (ProviderClient, error) {
	awsProviderConfig, ok := config.(*AWSProviderConfig)
	if !ok {
		return nil, fmt.Errorf("incorrect provider config provided")
	}
	if awsProviderConfig.Bucket == "" || awsProviderConfig.Region == "" {
		return nil, fmt.Errorf("incomplete AWS configuration")
	}
	awsSession, err := session.NewSession(&aws.Config{
		Region: &awsProviderConfig.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %v", err)
	}

	return &AWSProviderClient{
		Config: awsProviderConfig,
		Client: s3.New(awsSession),
	}, nil
}

// UploadFile uploads a file to S3 from an io.Reader.
func (client *AWSProviderClient) UploadFile(filePath string, reader io.Reader) error {
	uploader := s3manager.NewUploaderWithClient(client.Client)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(client.Config.Bucket),
		Key:    aws.String(filePath),
		Body:   reader,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %v", err)
	}
	return nil
}

// DeleteFile deletes a file from S3 by its key.
func (client *AWSProviderClient) DeleteFile(fileID string) error {
	_, err := client.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(client.Config.Bucket),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %v", err)
	}
	return nil
}

// ListFiles lists all files in a specified directory of the S3 bucket.
func (client *AWSProviderClient) ListFiles(directory string) ([]string, error) {
	var fileNames []string
	err := client.Client.ListObjectsV2Pages(&s3.ListObjectsV2Input{
		Bucket: aws.String(client.Config.Bucket),
		Prefix: aws.String(directory),
	}, func(page *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, obj := range page.Contents {
			fileNames = append(fileNames, *obj.Key)
		}
		return true // continue paging
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list files in S3: %v", err)
	}
	return fileNames, nil
}

// GetFile retrieves a file as an io.Reader by its key from S3.
func (client *AWSProviderClient) GetFile(fileID string) (io.Reader, error) {
	output, err := client.Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(client.Config.Bucket),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve file from S3: %v", err)
	}
	return output.Body, nil
}
