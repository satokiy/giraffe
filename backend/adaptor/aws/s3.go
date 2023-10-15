package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/samber/lo"
	"golang.org/x/xerrors"
)

type S3Config struct {
	Region string
	Bucket string
}

func NewS3Config(bucket, region string) S3Config {
	return S3Config{
		Region: region,
		Bucket: bucket,
	}
}

type S3Client interface {
	GetImageURLList() ([]string, error)
}

type S3ClientImpl struct {
	config S3Config
}

func NewS3Client(cfg S3Config) S3Client {
	return &S3ClientImpl{
		config: cfg,
	}
}

func (s3c *S3ClientImpl) GetImageURLList() ([]string, error) {

	return []string{"aaaa", "bbbb", "cccc"}, nil

	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(s3c.config.Region),
	}))

	svc := s3.New(session)
	result, err := svc.ListObjectsV2(
		&s3.ListObjectsV2Input{
			Bucket: aws.String(s3c.config.Bucket),
		})
	if err != nil {
		return nil, xerrors.Errorf("failed to S3 list object: %w", err)
	}

	objectList := lo.Map(result.Contents, func(obj *s3.Object, _ int) string {
		return *obj.Key
	})

	return objectList, nil
}
