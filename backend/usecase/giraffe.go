package usecase

import (
	"github.com/satokiy/giraffe/adaptor/repository"
	"github.com/satokiy/giraffe/model"
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
		return nil, err
	}

	// TODO
	image := imageList[0]

	return &model.GiraffeImage{URL: image}, nil
}
