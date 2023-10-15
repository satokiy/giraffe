package usecase

import (
	"math/rand"

	"github.com/satokiy/giraffe/adaptor/repository"
	"github.com/satokiy/giraffe/model"
	"golang.org/x/xerrors"
)

type GiraffeUsecase interface {
	GerImageRandom() (*model.GiraffeImage, error)
}

type GiraffeUsecaseImpl struct {
	giraffeRepository repository.GiraffeRepository
}

func NewGiraffeUsecaseImpl(r repository.GiraffeRepository) GiraffeUsecase {
	return &GiraffeUsecaseImpl{
		giraffeRepository: r,
	}
}

func (u *GiraffeUsecaseImpl) GerImageRandom() (*model.GiraffeImage, error) {
	imageList, err := u.giraffeRepository.GetGiraffeImageList()
	if err != nil {
		return nil, xerrors.Errorf("failed to get image list: %w", err)
	}

	// imageListからランダムに1つ取得
	imageKey := imageList[rand.Intn(len(imageList))]
	url, err := u.giraffeRepository.GetGiraffeImagePresignedURL(imageKey)
	if err != nil {
		return nil, xerrors.Errorf("failed to get presigned url: %w", err)
	}

	return &model.GiraffeImage{URL: url}, nil
}
