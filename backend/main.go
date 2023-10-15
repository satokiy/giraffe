package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/satokiy/giraffe/adaptor/aws"
	"github.com/satokiy/giraffe/adaptor/repository"
	"github.com/satokiy/giraffe/adaptor/restapi/controller"
	"github.com/satokiy/giraffe/router"
	"github.com/satokiy/giraffe/usecase"
)

func main() {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	s3cfg := aws.NewS3Config(
		os.Getenv("S3_BUCKET"),
		os.Getenv("AWS_REGION"),
	)

	s3Client := aws.NewS3Client(s3cfg)
	giraffeRepository := repository.NewGiraffeRepositoryImpl(s3Client)
	giraffeUsecase := usecase.NewGiraffeUsecaseImpl(giraffeRepository)
	giraffeController := controller.NewGiraffeControllerImpl(giraffeUsecase)
	e := router.NewRouter(giraffeController)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
