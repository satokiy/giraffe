package repository

import (
	"github.com/satokiy/giraffe/adaptor/aws"
	"golang.org/x/xerrors"
)

type GiraffeRepository interface {
	GetGiraffeImageList() ([]string, error)
}

type GiraffeRepositoryImpl struct {
	s3Client aws.S3Client
}

func NewGiraffeRepositoryImpl(s3c aws.S3Client) GiraffeRepository {
	return &GiraffeRepositoryImpl{
		s3Client: s3c,
	}
}

func (r *GiraffeRepositoryImpl) GetGiraffeImageList() ([]string, error) {
	objectList, err := r.s3Client.GetImageURLList()
	if err != nil {
		return nil, xerrors.Errorf("failed to get image url list: %w", err)
	}
	return objectList, nil
}
