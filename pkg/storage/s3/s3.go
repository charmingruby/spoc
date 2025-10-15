package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	client   *s3.Client
	uploader *manager.Uploader
	bucket   string
}

func New(bucket string) (*S3, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	return &S3{
		client:   client,
		uploader: uploader,
		bucket:   bucket,
	}, nil
}
