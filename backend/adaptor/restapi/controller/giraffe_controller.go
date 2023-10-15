package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satokiy/giraffe/adaptor/restapi/presenter"
	"github.com/satokiy/giraffe/usecase"
	"golang.org/x/xerrors"
)

type GiraffeController interface {
	GetRandom(c echo.Context) error
}

type GiraffeControllerImpl struct {
	gu usecase.GiraffeUsecase
}

func NewGiraffeControllerImpl(gu usecase.GiraffeUsecase) GiraffeController {
	return &GiraffeControllerImpl{gu}
}

func (gc *GiraffeControllerImpl) GetRandom(c echo.Context) error {
	giraffe, err := gc.gu.GerImageRandom()
	if err != nil {
		return xerrors.Errorf("failed to get random giraffe: %w", err)
	}

	response := presenter.GiraffeImageResponse(giraffe)

	// presigned urlの`&`がエンコードされないよう、カスタムエンコーダを実装する
	encoder := json.NewEncoder(c.Response().Writer)
	encoder.SetEscapeHTML(false)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return encoder.Encode(response)
}
