package s3

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (s *S3) Upload(ctx context.Context, data []byte) error {
	key := fmt.Sprintf("data/%s.json", time.Now().Format("20060102-150405"))

	if _, err := s.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/json"),
	}); err != nil {
		return fmt.Errorf("failed to upload to S3: %w", err)
	}

	return nil
}
