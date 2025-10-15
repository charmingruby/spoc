package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (s *S3) Reset(ctx context.Context) error {
	paginator := s3.NewListObjectsV2Paginator(s.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to list objects: %w", err)
		}

		if len(page.Contents) == 0 {
			continue
		}

		for i := 0; i < len(page.Contents); i += 1000 {
			end := min(i+1000, len(page.Contents))

			batch := page.Contents[i:end]
			if err := s.deleteBatchFromObjects(ctx, batch); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *S3) deleteBatchFromObjects(ctx context.Context, objects []s3types.Object) error {
	if len(objects) == 0 {
		return nil
	}

	objectIds := make([]s3types.ObjectIdentifier, len(objects))
	for i, obj := range objects {
		objectIds[i] = s3types.ObjectIdentifier{
			Key: obj.Key,
		}
	}

	_, err := s.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(s.bucket),
		Delete: &s3types.Delete{
			Objects: objectIds,
			Quiet:   aws.Bool(true),
		},
	})

	if err != nil {
		return fmt.Errorf("failed to delete batch: %w", err)
	}

	return nil
}
