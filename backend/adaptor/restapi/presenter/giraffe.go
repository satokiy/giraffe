package presenter

import (
	"github.com/satokiy/giraffe/adaptor/restapi/apiresponse"
	"github.com/satokiy/giraffe/model"
)

func GiraffeImageResponse(m *model.GiraffeImage) *apiresponse.GiraffeResponse {
	return &apiresponse.GiraffeResponse{
		URL: m.URL,
	}
}
